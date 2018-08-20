package job

import (
	"context"
	"encoding/gob"
	"time"

	"github.com/stretchr/testify/require"
)

type simpleJob struct {
	Data string
}

func (a *simpleJob) Run(ctx context.Context, r Runtime) error {
	r.Log().Info("simpleJob running!")
	return nil
}

func init() {
	gob.Register(&simpleJob{})
}

// TestCompatibility checks store functionality.
func TestCompatibility(t require.TestingT, store Store) {
	require := require.New(t)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := store.Delete(ctx, "")
	require.Error(err)

	_, err = store.Load(ctx, "")
	require.Error(err)

	run := &simpleJob{}
	rec, err := store.Create(ctx, run, "test")
	require.NoError(err)
	require.NotNil(rec)
	require.NotZero(rec.JobID())

	recIDs, err := store.SearchByTag(ctx, "test")
	require.NoError(err)
	require.Len(recIDs, 1)

	first, err := rec.LastSnapshot(ctx)
	require.NoError(err)
	require.NotNil(first)
	require.False(first.IsFinal())
	require.NoError(first.Result())

	snaps, err := rec.SnapshotsSince(ctx, first)
	require.NoError(err)
	require.Len(snaps, 0)

	snaps, err = rec.SnapshotsSince(ctx, nil)
	require.NoError(err)
	require.Len(snaps, 1)

	run.Data = "snap1"
	err = rec.AddSnapshot(ctx, run, WithDescription("snap1"))
	require.NoError(err)

	snap1, err := rec.LastSnapshot(ctx)
	require.NoError(err)
	waitResult := make(chan error)
	go func() {
		waitResult <- rec.WaitForUpdate(ctx, snap1)
	}()

	run.Data = "final"
	err = rec.AddSnapshot(ctx, run, AsFinal(true), WithDescription("final"))
	require.NoError(err)

	err = <-waitResult
	require.NoError(err)

	snaps, err = rec.SnapshotsSince(ctx, first)
	require.NoError(err)
	require.Len(snaps, 2)

	last, err := rec.LastSnapshot(ctx)
	require.NoError(err)
	require.NotNil(last)
	require.True(last.IsFinal())
	require.NoError(last.Result())

	run2, err := snaps[0].Runnable()
	require.NoError(err)
	require.Equal("snap1", run2.(*simpleJob).Data)

	run3, err := last.Runnable()
	require.NoError(err)
	require.Equal("final", run3.(*simpleJob).Data)

	err = store.Delete(ctx, rec.JobID())
	require.NoError(err)
}
