package job

import (
	"context"
	"io"
	"time"
)

// Store is an interface for managing job state.
type Store interface {
	io.Closer

	// Create a new job record with an initial state snapshot.
	// The initial snapshot contains the runnable state.
	Create(ctx context.Context, run Runnable, tags ...string) (Record, error)

	// Load returns an existing state record, or an error.
	// Returns ErrNotFound if the record does not exist.
	Load(ctx context.Context, id string) (Record, error)

	// SearchByTag returns job ids associated with the given tag.
	SearchByTag(ctx context.Context, tag string) (ids []string, err error)

	// Delete removes the specified job, or returns an error.
	// Returns ErrNotFound if the record does not exist.
	Delete(ctx context.Context, id string) error
}

// Record is an interface to a stored job.
// Operations are not guaranteed to be thread safe.
type Record interface {
	// JobID is the unique id of the job.
	JobID() string

	// Tags returns the tags specified in job creation.
	Tags() []string

	// AddSnapshot creates a new snapshot for the job.
	AddSnapshot(ctx context.Context, run Runnable, opts ...SnapshotOpt) error

	// LastSnapshot returns the latest snapshot for the record.
	// All records have at a snapshot, so snapshot is never nil unless an
	// error occurs.
	LastSnapshot(ctx context.Context) (Snapshot, error)

	// SnapshotsSince returns snapshots which come after the given snapshot,
	// or all available snapshots if the after parameter is nil.
	SnapshotsSince(ctx context.Context, after Snapshot) ([]Snapshot, error)

	// WaitForUpdate blocks until there has been an update to the record which
	// occurs after the given snapshot.  This call may return immediately if an
	// update already exists.
	// The caller is responsible for loading the new snapshot.
	WaitForUpdate(ctx context.Context, snap Snapshot) error
}

// Snapshot provides access to job state at a point in time.
type Snapshot interface {
	// Runnable returns the runnable object from the job.
	Runnable() (Runnable, error)

	// Description contains arbitrary text associated with the snapshot.
	Description() string

	// IsFinal returns true after job has been finalized.
	IsFinal() bool

	// Result is the last error code stored with the job.
	Result() error

	// When contains the wall-clock time that the snapshot was added.
	When() time.Time
}

// SnapshotSetter is an API for use in constructing checkpoints.
type SnapshotSetter interface {
	SetError(err error)
	SetDescription(desc string)
	SetFinal(final bool)
}

// SnapshotOpt is an option when taking a checkpoint.
type SnapshotOpt func(SnapshotSetter)

// AsFinal marks a checkpoint as final.
func AsFinal(final bool) SnapshotOpt {
	return func(c SnapshotSetter) {
		c.SetFinal(final)
	}
}

// WithError indicates that the checkpoint failed.
func WithError(err error) SnapshotOpt {
	return func(c SnapshotSetter) {
		c.SetError(err)
	}
}

// WithDescription provides a description for the checkpoint.
func WithDescription(desc string) SnapshotOpt {
	return func(c SnapshotSetter) {
		c.SetDescription(desc)
	}
}
