package workflow

import (
	"context"

	log "github.com/sirupsen/logrus"
)

// Workflow is an instance of a workflow.
//
// Parameters should be part of the workflow object, and must be serializable
// for deferred execution.  Register workflow implementation with gob.Register.
type Workflow interface {
	// Result returns an error if the job failed, or nil if the job succeeded.
	Run(ctx context.Context, r Runtime)
}

// NewJob wraps a workflow as a runnable job. Use this to schedule and run
// workflows with job.Scheduler.
//
// Note that the type passed as Workflow must be registered with encoding/gob.
// Call "gob.Register(&MyType{})" in an init function to register the concrete
// types.
func NewJob(name string, wf Workflow) *JobWrapper {
	return &JobWrapper{
		Name:     name,
		Workflow: wf,
	}
}

// OpFN is shorthand for a workflow operation function.
type OpFN func()

// Runtime provides access to runtime operations for workflows.
// This is provided to the workflow instance when it is run.
type Runtime interface {
	// Log returns a logging channel for the current stage.
	Log() log.FieldLogger

	// ID returns the unique job id.
	ID() string

	// Context returns the current context.
	Context() context.Context

	// Do adds a function which can be executed multiple times.
	//
	// The function may be attempted multiple times. The operation should be
	// idempotent, and result in the same state over multiple calls.
	Do(fn OpFN, opts ...Opt)

	// Embed add a new workflow as a step in the current workflow.
	Embed(wf Workflow, opts ...Opt)

	// OnRollback registers a rollback operation, which will only be executed if
	// the workflow is rolled back beyond the rollback point.
	//
	// Place a rollback before an operation if the rollback is stateless and can
	// execute even if the success or failure of the stage is uncertain.  Place
	// a rollback after an operation if the rollback depends on state of
	// previous steps and can only run if the operation is certain to have
	// succeeded.
	OnRollback(undo OpFN)

	// Fail the workflow execution step. If the error is an un-retryable failure,
	// then retries are not allowed, otherwise the workflow may be retried if
	// retry policy allows.
	//
	// When a final failure occurs, either because the failure is un-retryable,
	// or because retries have been exhausted, then undo operations are run to
	// revert stages that did succeed (or may have succeeded).
	Fail(err error)
}

type runOpts struct {
	vars        []interface{}
	rollbackAll bool
	name        string
}

// Opt is a workflow execution option.
type Opt func(*runOpts)

// State marks variable references as having been mutated by the
// stage.  This ensures that the values are stored as part of the stage's cached
// state, and recovered when using cached values in a retry.
func State(vars ...interface{}) Opt {
	return func(c *runOpts) {
		c.vars = append(c.vars, vars...)
	}
}

// RollbackOnFailure indicates that if the stage fails, then the entire
// workflow must be rolled back.  The workflow could still be retried, but not
// using cached results from earlier stages.
//
// This is useful when implementing optimistic locking, where a late-stage
// commit failure requires restarting the workflow.
func RollbackOnFailure() Opt {
	return func(c *runOpts) {
		c.rollbackAll = true
	}
}

// Name sets a name for the workflow stage.
// If not provided, then a generated name is used.
func Name(name string) Opt {
	return func(c *runOpts) {
		c.name = name
	}
}
