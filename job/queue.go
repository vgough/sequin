package job

import (
	"context"
	"io"
)

// Work provides details about a work item obtained from a Queue.
type Work interface {
	ID() string

	// Acknowledge marks a work item as complete, so that it will not be
	// rescheduled.
	//
	// If the work is not acknowledged when the handler completes, then it can
	// be rescheduled on another worker.
	Acknowledge() error
}

// Queue is used to schedule jobs and find jobs to run.
type Queue interface {
	io.Closer

	// Schedule a job for execution.
	Schedule(ctx context.Context, id string) error

	// RunNext waits for and processes the next available job.
	// The callback function is given information about the work item.
	RunNext(ctx context.Context, f func(w Work) error) error
}
