package sql

import (
	"bytes"
	"encoding/gob"
	"time"

	"github.com/pkg/errors"
	"github.com/vgough/sequin/job"
)

// Snapshot holds the result of a job execution.
type Snapshot struct {
	ID uint `gorm:"primary_key"`

	StoredJobID uint      `gorm:"index:idx_job_created"`
	CreatedAt   time.Time `gorm:"index:idx_job_created"`

	// Desc is a human-readable description of the snapshot.
	Desc string

	// Final is true if the job completed, either successfully or with an
	// error.
	// Once a job is marked final, no further updates are allowed.
	Final bool

	// Error stores any runtime error during execution.
	Error string

	// State holds a serialized job state.
	State []byte
}

// Runnable returns the job runnable.
func (s *Snapshot) Runnable() (job.Runnable, error) {
	buf := bytes.NewBuffer(s.State)
	dec := gob.NewDecoder(buf)
	var decoded job.Runnable
	if err := dec.Decode(&decoded); err != nil {
		return nil, errors.Wrapf(err, "unable to decode snapshot %q", s.ID)
	}
	return decoded, nil
}

// Description returns the snapshot description.
func (s *Snapshot) Description() string {
	return s.Desc
}

// IsFinal returns true when the job has been finalized.
func (s *Snapshot) IsFinal() bool {
	return s.Final
}

// Result contains the final result of the job.
func (s *Snapshot) Result() error {
	if s.Error == "" {
		return nil
	}
	return errors.New(s.Error)
}

// When returns the time the snapshot was created.
func (s *Snapshot) When() time.Time {
	return s.CreatedAt
}

// SetError sets the error result of the snapshot.
func (s *Snapshot) SetError(err error) {
	if err == nil {
		s.Error = ""
	} else {
		s.Error = err.Error()
	}
}

// SetDescription changes the snapshot description.
func (s *Snapshot) SetDescription(desc string) {
	s.Desc = desc
}

// SetFinal sets the final tag on a snapshot.
func (s *Snapshot) SetFinal(f bool) {
	s.Final = f
}
