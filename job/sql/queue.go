package sql

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"github.com/vgough/sequin/job"
)

var chLock sync.Mutex
var channels = map[string]chan *WorkItem{}
var pollInterval = 1 * time.Second
var updateInterval = 15 * time.Second
var lockTime = updateInterval*2 + pollInterval

// Queue implements job.Queue in a SQL store.
type Queue struct {
	db      *gorm.DB
	channel string
	local   chan *WorkItem
}

var _ job.Queue = &Queue{}

// NewQueue returns a new job scheduling queue.
func NewQueue(db *gorm.DB, channelName string) *Queue {
	chLock.Lock()
	ch, ok := channels[channelName]
	if !ok {
		ch = make(chan *WorkItem)
		channels[channelName] = ch
	}
	chLock.Unlock()

	return &Queue{db: db, channel: channelName, local: ch}
}

// TestQueue returns a queue suitable for unit tests.
func TestQueue(t Testing) *Queue {
	id := uuid.NewV4()
	return NewQueue(getTestDB(t), fmt.Sprintf("test-%s", id.String()))
}

// Close releases resources.
func (q *Queue) Close() error {
	chLock.Lock()
	delete(channels, q.channel)
	chLock.Unlock()
	return nil
}

// Schedule makes a work item executable.
func (q *Queue) Schedule(ctx context.Context, id string) error {
	nid, err := parseID(id)
	if err != nil {
		return err
	}

	it := &WorkItem{
		StoredJobID: nid,
		Channel:     q.channel,
		RunnableAt:  time.Now(), // TODO
		db:          q.db,
	}
	if err := q.db.Create(it).Error; err != nil {
		return errors.Wrapf(err, "unable to schedule job %s", id)
	}
	ll := log.With().Str("id", id).Logger()
	ll.Info().Msg("scheduled work item")

	// Optimistically attempt to have local worker handle it, if one exists.
	select {
	case q.local <- it:
		ll.Info().Msg("passed job to local worker")
	default: // non-blocking, leave it
	}
	return nil
}

func (q *Queue) takeLock(it *WorkItem) error {
	log.Info().Uint("id", it.StoredJobID).Msg("locking work item")
	it.RunnableAt = time.Now().Add(lockTime)
	oldCount := it.AttemptCount
	it.AttemptCount++
	return q.db.Model(it).
		Where("attempt_count = ?", oldCount).
		Update(map[string]interface{}{
			"runnable_at":   it.RunnableAt,
			"attempt_count": it.AttemptCount,
		}).Error

}

func (q *Queue) extendLock(it *WorkItem) error {
	log.Info().Uint("id", it.StoredJobID).Msg("extending work item")
	it.RunnableAt = time.Now().Add(lockTime)
	return q.db.Model(it).
		Where("attempt_count = ?", it.AttemptCount).
		Update("runnable_at", it.RunnableAt).Error
}

// RunNext runs the next available work item.
func (q *Queue) RunNext(ctx context.Context, f func(w job.Work) error) error {
	for {
		var it *WorkItem
		select {
		case <-ctx.Done():
			return ctx.Err()
		case it = <-q.local:
			// Acquired work from local queue.
		case <-time.After(pollInterval):
			// Look for available work in DB.
			var item WorkItem
			now := time.Now()
			if err := q.db.Where("channel = ?", q.channel).
				Where("runnable_at < ?", now).
				Order("runnable_at").
				First(&item).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					log.Info().Err(err).Msg("unable to find runnable item")
				}
				continue
			}
			it = &item
		}

		ll := log.With().Uint("id", it.StoredJobID).Logger()
		ll.Info().Msg("obtained work item")

		// Update item to acquire it.
		if err := q.takeLock(it); err != nil {
			log.Warn().Err(err).Msg("unable to lock item")
			continue
		}

		// Start processing.
		done := make(chan error, 1)
		go func() {
			ll.Info().Msg("starting job")
			defer func() { ll.Info().Msg("job completed") }()
			done <- f(it)
		}()

		// Keep updating record while working to keep work from expiring.
		for {
			select {
			case err := <-done:
				if (err == nil) || job.IsUnrecoverable(err) {
					return err
				}

				// TODO: retry logic.
				return errors.Wrapf(err, "failed to process %q", it.StoredJobID)

			case <-time.After(updateInterval):
				if err := q.extendLock(it); err != nil {
					log.Warn().Err(err).Msg("unable to lock item")
				}
			}
		}
	}
}

// WorkItem is a table for holding scheduled work.
type WorkItem struct {
	StoredJobID uint `gorm:"unique_index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Channel    string    `gorm:"index:idx_channel_runtime"`
	RunnableAt time.Time `gorm:"index:idx_channel_runtime"`

	AttemptCount int

	db *gorm.DB
}

// ID returns the unique ID of the work item.
func (w *WorkItem) ID() string {
	return strconv.FormatUint(uint64(w.StoredJobID), 10)
}

// Acknowledge marks a work item as completed.
func (w *WorkItem) Acknowledge() error {
	log.Info().Uint("id", w.StoredJobID).Msg("deleting work item")
	return w.db.Where("stored_job_id = ?", w.StoredJobID).Delete(w).Error
}
