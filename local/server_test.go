package local

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vgough/sequin"
)

func TestServer(t *testing.T) {
	s := NewServer()
	ctx := sequin.WithRuntime(context.Background(), s)

	ok, err := IsEven(ctx, 0)
	require.NoError(t, err)
	require.True(t, ok)

	sum, err := StatefulOp(ctx, 1)
	require.ErrorContains(t, err, "test error")
	require.Equal(t, 1, sum)

	sum, err = StatefulOp(ctx, 1)
	require.NoError(t, err)
	require.Equal(t, 2, sum)
}

var IsEven = sequin.Register(isEven)

func isEven(ctx context.Context, in int) (bool, error) {
	if in < 0 {
		return false, errors.New("negative values not supported")
	}
	if in%2 == 0 {
		return true, nil
	}
	return false, nil
}

var StatefulOp = sequin.Register(statefulOp)

func statefulOp(ctx context.Context, iterations int) (int, error) {
	var numRuns int
	if err := sequin.OpInit(ctx, sequin.Persist("runCount", &numRuns)); err != nil {
		return 0, err
	}
	numRuns++
	if err := sequin.Checkpoint(ctx); err != nil {
		return 0, err
	}

	if numRuns < 2 {
		return numRuns, errors.New("test error")
	}
	return numRuns, nil
}
