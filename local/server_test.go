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

	_, err = StateTest(ctx)
	require.NoError(t, err)

	oi, of, ob, os, os2, err := ParamTests(ctx, 1, 2.0, true, "test",
		SimpleStruct{Float64State: 3.0})
	require.NoError(t, err)
	require.Equal(t, 1, oi)
	require.Equal(t, 2.0, of)
	require.True(t, ob)
	require.Equal(t, "test", os)
	require.Equal(t, SimpleStruct{Float64State: 3.0}, os2)
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

var StateTest = sequin.Register(stateTest)

type SimpleStruct struct {
	Float64State float64
}

func stateTest(ctx context.Context) (int, error) {
	var intState int
	var strState string
	var structState SimpleStruct
	if err := sequin.OpInit(ctx,
		sequin.Persist("intState", &intState),
		sequin.Persist("strState", &strState),
		sequin.Persist("structState", &structState)); err != nil {
		return 0, err
	}
	intState++
	strState += "a"
	structState.Float64State += 1.0
	if err := sequin.Checkpoint(ctx); err != nil {
		return 0, err
	}
	return 0, nil
}

var ParamTests = sequin.Register(paramTests)

func paramTests(ctx context.Context, ini int, inf float64,
	inb bool, instr string, inStruct SimpleStruct) (int, float64, bool, string, SimpleStruct, error) {
	return ini, inf, inb, instr, inStruct, nil
}
