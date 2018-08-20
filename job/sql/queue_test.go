package sql

import (
	"context"
	"encoding/gob"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vgough/sequin/job"
)

type simpleJob struct {
	Result   error
	RunCount int
}

func (a *simpleJob) Run(ctx context.Context, r job.Runtime) error {
	r.Log().Info("simpleJob running!")
	a.RunCount++
	return a.Result
}

func init() {
	gob.Register(&simpleJob{})
}

func TestQueuing(t *testing.T) {
	store := TestStore(t)
	q := TestQueue(t)

	sch := job.NewScheduler(q, store)
	defer sch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go sch.WorkerLoop(ctx)

	t.Run("SimpleJob", func(t *testing.T) {
		job := &simpleJob{}
		rec, err := sch.Enqueue(ctx, job, "name=queue.SimpleJob")
		require.NoError(t, err)

		snap, err := sch.WaitForCompletion(ctx, rec)
		require.NoError(t, err)
		require.True(t, snap.IsFinal())
		require.NoError(t, snap.Result())

		store.Delete(ctx, rec.JobID())
	})
}
