package job

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"
)

// Runtime is the scheduling API presented to jobs at runtime.
type Runtime interface {
	// Update the state of the stored job.
	// It is not required for jobs to call this unless they need to update state
	// during a run.  When a job completes, the state is always updated.
	Update(ctx context.Context, opts ...SnapshotOpt) error

	// ID returns the unique job id.
	ID() string
}

// Run processes a job record.
func Run(ctx context.Context, rec Record) error {
	ctx, span := trace.StartSpan(ctx, "job.Run")
	defer span.End()
	span.AddAttributes(trace.StringAttribute("job_id", rec.JobID()))

	ll := log.With().Str("id", rec.JobID()).Logger()

	snap, err := rec.LastSnapshot(ctx)
	if err != nil {
		return err
	}
	if snap.IsFinal() {
		return snap.Result()
	}

	run, err := snap.Runnable()
	if err != nil {
		return err
	}

	rt := &runtime{
		rec: rec,
		run: run,
	}
	ll.Info().Msg("starting job")
	err = run.Run(ctx, rt)

	var final bool
	if _, unrec := err.(UnrecoverableError); unrec || err == nil {
		ll.Info().Msg("finalizing job")
		final = true
	}

	if uerr := rec.AddSnapshot(ctx, run, WithError(err), AsFinal(final)); uerr != nil {
		return errors.Wrap(uerr, "unable to update job state")
	}

	return nil
}

// runtime implements the Runtime API for the sql store.
type runtime struct {
	rec Record
	run Runnable
}

var _ Runtime = &runtime{}

// ID returns the job id.
func (s *runtime) ID() string {
	return s.rec.JobID()
}

// Update implements Runtime.
func (s *runtime) Update(ctx context.Context, opts ...SnapshotOpt) error {
	return s.rec.AddSnapshot(ctx, s.run, opts...)
}
