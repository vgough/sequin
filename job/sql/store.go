package sql

import (
	"bytes"
	"context"
	"encoding/gob"
	"strconv"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/karlseguin/rcache"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/vgough/sequin/job"
	"go.opencensus.io/trace"
)

// Testing is the minimum amount of testing.T needed for running tests.
// This avoid pulling in the testing package during normal builds.
type Testing interface {
	Logf(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

// Store is a database backend for storing job state.
type Store struct {
	db *gorm.DB

	updateListeners *rcache.Cache
}

var _ job.Store = &Store{}

var testDBLocation = "/tmp/workflow.db"

// var testDBLocation = "file::memory:?mode=memory&cache=shared"
var testDB *gorm.DB
var testDBInit sync.Once

// NewStore returns a new storage backend.
func NewStore(db *gorm.DB) *Store {
	ch := func(key string) interface{} {
		log.Info().Str("key", key).Msg("creating update channel")
		return make(chan struct{})
	}
	return &Store{
		db:              db,
		updateListeners: rcache.New(ch, 5*pollInterval),
	}
}

func getTestDB(t Testing) *gorm.DB {
	testDBInit.Do(func() {
		log.Info().Msg("constructing global test db")
		db, err := gorm.Open("sqlite3", testDBLocation)
		if err != nil {
			t.Fatalf("unable to create DB: %s", err)
		}
		// db.LogMode(true)
		db.BlockGlobalUpdate(true)
		if err := MigrateDB(db); err != nil {
			t.Fatalf("unable to initialize DB: %s", err)
		}
		testDB = db
	})
	return testDB
}

// TestStore returns a new State instance for use in testing.
// Note, you much load sqlite dialect support for this to work.
// This can be done by including the sqlite dialect package:
//    _ "github.com/jinzhu/gorm/dialects/sqlite" // For test db.
func TestStore(t Testing) *Store {
	return NewStore(getTestDB(t))
}

// MigrateDB constructs the database scheme if it does not exist.  It will also
// add columns if necessary, but will not remove existing columns or indexes.
// This isn't necessarily a good idea in production, but is useful for tests
// which typically start with an uninitialized DB.
func MigrateDB(db *gorm.DB) error {
	db.AutoMigrate(&StoredJob{}, &Snapshot{}, &Tag{}, &WorkItem{})
	return db.Error
}

func parseID(id string) (uint, error) {
	nid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "invalid id %q", id)
	}
	return uint(nid), nil
}

// Close releases resources.
func (s *Store) Close() error {
	return nil
}

// Create adds a new job to stored_jobs.
func (s *Store) Create(ctx context.Context, j job.Runnable, tags ...string) (job.Record, error) {
	ctx, span := trace.StartSpan(ctx, "record.create")
	defer span.End()

	data, err := encodeJob(j)
	if err != nil {
		return nil, err
	}

	var tagList []Tag
	for _, t := range tags {
		tagList = append(tagList, Tag{
			Value: t,
		})
	}

	snap := &Snapshot{
		State: data,
	}
	job := &StoredJob{
		JobTags:   tagList,
		Snapshots: []*Snapshot{snap},
		store:     s,
	}

	if err := s.db.Create(job).Error; err != nil {
		return nil, errors.Wrap(err, "unable to create job")
	}

	span.AddAttributes(trace.StringAttribute("job_id", job.JobID()))
	return job, nil
}

// Load implements Scheduler.
func (s *Store) Load(ctx context.Context, id string) (job.Record, error) {
	ctx, span := trace.StartSpan(ctx, "record.load")
	defer span.End()
	span.AddAttributes(trace.StringAttribute("job_id", id))

	nid, err := parseID(id)
	if err != nil {
		return nil, err
	}

	job := StoredJob{
		ID:    nid,
		store: s,
	}
	if err := s.db.First(&job).Error; err != nil {
		return nil, errors.Wrapf(err, "unable to lookup job %q", id)
	}

	if err := s.db.Model(&job).Related(&job.JobTags).Error; err != nil {
		return nil, errors.Wrapf(err, "unable to load tags for job %q", id)
	}

	return &job, nil
}

// SearchByTag returns job ids associated with the given tag.
func (s *Store) SearchByTag(ctx context.Context, tag string) ([]string, error) {
	var tags []Tag
	if err := s.db.Where("value = ?", tag).Find(&tags).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	ids := make([]string, len(tags))
	var job StoredJob
	for i, t := range tags {
		job.ID = t.StoredJobID
		ids[i] = job.JobID()
	}

	return ids, nil
}

// Delete removes records for a job.
func (s *Store) Delete(ctx context.Context, id string) error {
	ctx, span := trace.StartSpan(ctx, "record.delete")
	defer span.End()
	span.AddAttributes(trace.StringAttribute("job_id", id))

	nid, err := parseID(id)
	if err != nil {
		return err
	}

	if err := s.db.Delete(&StoredJob{ID: nid}).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	if err := s.db.Where("stored_job_id = ?", nid).Delete(&Tag{}).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	if err := s.db.Where("stored_job_id = ?", nid).Delete(&Snapshot{}).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	return nil
}

func (s *Store) updateChannel(id uint) <-chan struct{} {
	return s.updateListeners.Get(strconv.FormatUint(uint64(id), 10)).(chan struct{})
}

// notifyLocal wakes up any local threads waiting for a job to be updated.
func (s *Store) notifyLocal(id uint) {
	key := strconv.FormatUint(uint64(id), 10)
	ch := s.updateListeners.Get(key).(chan struct{})
	s.updateListeners.Delete(key)
	close(ch)
}

func encodeJob(r job.Runnable) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(&r); err != nil {
		return nil, errors.Wrap(err, "unable to encode job")
	}

	return buf.Bytes(), nil
}
