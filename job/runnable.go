package job

import "context"

// Runnable is a generic interface for a job that can be scheduled with
// Scheduler. When the job is run, it is given a context.
type Runnable interface {
	// Run returns an error if the job failed.  If the error is of type
	// Unrecoverable, then the job cannot be retried.
	Run(ctx context.Context, r Runtime) error
}
