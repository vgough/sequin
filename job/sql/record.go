package sql

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/vgough/sequin/job"
	"go.opencensus.io/trace"
)

// StoredJob contains persistent storage for a job.
type StoredJob struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time

	// Tags contains the tags associated with the job.
	JobTags []Tag

	Snapshots []*Snapshot

	store *Store `gorm:"-"`
}

var _ job.Record = &StoredJob{}

// Tag contains an indexed value and associated StoredJob foreign key.
type Tag struct {
	StoredJobID uint   `gorm:"primary_key"`
	Value       string `gorm:"primary_key"`
}

// JobID returns the unique id of the job.
func (j *StoredJob) JobID() string {
	return strconv.FormatUint(uint64(j.ID), 10)
}

// Tags returns tags in string form.
func (j *StoredJob) Tags() []string {
	tags := make([]string, len(j.JobTags))
	for i, t := range j.JobTags {
		tags[i] = t.Value
	}
	return tags
}

// LastSnapshot returns the last available snapshot, or an error.
func (j *StoredJob) LastSnapshot(ctx context.Context) (job.Snapshot, error) {
	ctx, span := trace.StartSpan(ctx, "snapshot.last")
	defer span.End()
	span.AddAttributes(trace.StringAttribute("job_id", j.JobID()))

	var snap Snapshot
	if err := j.store.db.Where("stored_job_id = ?", j.ID).
		Order("created_at DESC").
		First(&snap).Error; err != nil {
		return nil, err
	}

	return &snap, nil
}

// SnapshotsSince returns all snapshots after the given snapshot.
func (j *StoredJob) SnapshotsSince(ctx context.Context, after job.Snapshot) ([]job.Snapshot, error) {
	ctx, span := trace.StartSpan(ctx, "snapshot.since")
	defer span.End()
	span.AddAttributes(trace.StringAttribute("job_id", j.JobID()))

	var snaps []*Snapshot
	db := j.store.db.Where("stored_job_id = ?", j.ID)
	if after != nil {
		db = db.Where("created_at > ?", after.When())
	}

	if err := db.Find(&snaps).Error; err != nil {
		return nil, err
	}
	result := make([]job.Snapshot, len(snaps))
	for i, s := range snaps {
		result[i] = s
	}
	return result, nil
}

// AddSnapshot stores new runtime state.
func (j *StoredJob) AddSnapshot(ctx context.Context, run job.Runnable, opts ...job.SnapshotOpt) error {
	ctx, span := trace.StartSpan(ctx, "snapshot.create")
	defer span.End()
	span.AddAttributes(trace.StringAttribute("job_id", j.JobID()))

	data, err := encodeJob(run)
	if err != nil {
		return err
	}

	snap := &Snapshot{
		StoredJobID: j.ID,
		State:       data,
	}
	for _, o := range opts {
		o(snap)
	}

	if snap.Description() != "" {
		span.AddAttributes(trace.StringAttribute("description", snap.Description()))
	}

	if err := j.store.db.Create(snap).Error; err != nil {
		return errors.Wrapf(err, "unable to create checkpoint for %d", j.ID)
	}

	log.WithFields(log.Fields{
		"id":          j.ID,
		"snapshot_id": snap.ID,
	}).Info("added state snapshot")
	j.store.notifyLocal(j.ID)

	return nil
}

// WaitForUpdate waits for an update to occur.
func (j *StoredJob) WaitForUpdate(ctx context.Context, snap job.Snapshot) error {
	ctx, span := trace.StartSpan(ctx, "snapshot.wait")
	defer span.End()
	span.AddAttributes(trace.StringAttribute("job_id", j.JobID()))

	if snap.IsFinal() {
		return errors.New("job is final, no updates will occur")
	}

	ll := log.WithField("id", j.ID)
	for {
		var count int
		ll.Info("checking for newer snapshots")
		updated := j.store.updateChannel(j.ID)

		if err := j.store.db.Model(&Snapshot{}).
			Where("stored_job_id = ?", j.ID).
			Where("created_at > ?", snap.When()).
			Count(&count).Error; err != nil {
			return err
		}

		span.Annotate([]trace.Attribute{
			trace.Int64Attribute("update_count", int64(count)),
		}, "snapshot query")
		if count > 0 {
			return nil
		}

		select {
		case <-ctx.Done():
			span.Annotate([]trace.Attribute{}, "context done")
			return ctx.Err()
		case <-updated:
			span.Annotate([]trace.Attribute{}, "local notify")
		case <-time.After(pollInterval):
			span.Annotate([]trace.Attribute{}, "poll timeout")
		}
	}
}
