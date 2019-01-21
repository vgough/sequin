package job

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"
)

// Scheduler is a helper for constructing and executing jobs.
type Scheduler struct {
	q     Queue
	store Store
	wg    sync.WaitGroup // for tracking background tasks
}

// NewScheduler creates a new scheduler.
func NewScheduler(q Queue, store Store) *Scheduler {
	return &Scheduler{
		q:     q,
		store: store,
	}
}

// Close blocks until background loops terminate before calling close on
// the store and queue.
// It is critical that contexts used in the loops are canceled, otherwise
// Close may block indefinitely.
func (s *Scheduler) Close() error {
	s.wg.Wait()
	s.q.Close()
	s.store.Close()
	return nil
}

// Enqueue records and runs a job.
func (s *Scheduler) Enqueue(ctx context.Context, j Runnable, tags ...string) (Record, error) {
	ctx, span := trace.StartSpan(ctx, "enqueue")
	defer span.End()

	rec, err := s.store.Create(ctx, j, tags...)
	if err != nil {
		return nil, err
	}

	if err := s.q.Schedule(ctx, rec.JobID()); err != nil {
		return nil, err
	}
	return rec, nil
}

func (s *Scheduler) runFN(ctx context.Context) func(w Work) error {
	return func(w Work) error {
		ll := log.With().Str("id", w.ID()).Logger()

		rec, err := s.store.Load(ctx, w.ID())
		if err != nil {
			return err
		}

		if err := Run(ctx, rec); err != nil {
			return err
		}

		ll.Info().Msg("checking job state")
		// If the job is finalized, then ack the work item.
		if snap, err := rec.LastSnapshot(ctx); err != nil {
			ll.Error().Err(err).Msg("unable to load last snapshot")
		} else if snap.IsFinal() {
			if err := w.Acknowledge(); err != nil {
				ll.Warn().Err(err).Msg("unable to ack message")
			}
		} else {
			log.Info().Interface("rec", rec).Msg("job incomplete")
		}

		return err
	}
}

// WorkerLoop runs jobs in a loop.  This does not return unless an internal
// failure occurs or the context is canceled.
func (s *Scheduler) WorkerLoop(ctx context.Context) error {
	s.wg.Add(1)
	defer s.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		sctx, span := trace.StartSpan(ctx, "worker")
		if err := s.q.RunNext(sctx, s.runFN(sctx)); err != nil {
			log.Debug().Err(err).Msg("job failed")
		}
		span.End()
	}
}

// WaitForCompletion waits until the given job is marked as completed.
func (s *Scheduler) WaitForCompletion(ctx context.Context, rec Record) (Snapshot, error) {
	s.wg.Add(1)
	defer s.wg.Done()

	ctx, span := trace.StartSpan(ctx, "wait")
	defer span.End()

	ll := log.With().Str("id", rec.JobID()).Logger()
	for {
		ll.Info().Msg("checking completion state")
		snap, err := rec.LastSnapshot(ctx)
		if err != nil {
			return nil, err
		}

		if snap.IsFinal() {
			err := snap.Result()
			ll.Info().Err(err).Msg("waitForCompletion, job finalized")
			return snap, nil
		}

		if err := rec.WaitForUpdate(ctx, snap); err != nil {
			log.Error().Err(err).Msg("waitForCompletion failing")
			return nil, err
		}
	}
}
