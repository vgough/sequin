package sequin

import (
	"context"
	"errors"
	"log/slog"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vgough/sequin/registry"
)

func TestRegister(t *testing.T) {
	ok, err := IsEven(context.Background(), 0)
	require.NoError(t, err)
	require.True(t, ok)

	endpoints := registry.RegisteredEndpoints()
	t.Log(endpoints)
	require.Contains(t, endpoints, "github.com/arg0net/stateful.isEven")
}

func TestRuntime(t *testing.T) {
	ctx := context.Background()
	var callCount int
	rt := &MockRuntime{
		ExecFn: func(ep *registry.Endpoint, args []reflect.Value) []reflect.Value {
			callCount++
			return ep.Exec(args)
		},
	}
	ctx = WithRuntime(ctx, rt)

	ok, err := IsEven(ctx, 0)
	require.NoError(t, err)
	require.True(t, ok)
	require.Equal(t, 1, callCount)
}

func TestRuntimeError(t *testing.T) {
	ctx := context.Background()
	rt := &MockRuntime{
		ExecFn: func(ep *registry.Endpoint, args []reflect.Value) []reflect.Value {
			slog.Info("mock error", "outputTypes", ep.OutputTypes)
			return ep.MakeError(errors.New("mock error"))
		},
	}
	ctx = WithRuntime(ctx, rt)

	_, err := IsEven(ctx, 0)
	require.Error(t, err)
	require.Contains(t, err.Error(), "mock error")
}

var IsEven = Register(isEven)

func isEven(ctx context.Context, in int) (bool, error) {
	if in < 0 {
		return false, errors.New("negative values not supported")
	}
	if in%2 == 0 {
		return true, nil
	}
	return false, nil
}

type MockRuntime struct {
	ExecFn func(ep *registry.Endpoint, args []reflect.Value) []reflect.Value
}

func (rt *MockRuntime) Exec(ep *registry.Endpoint, args []reflect.Value) []reflect.Value {
	return rt.ExecFn(ep, args)
}

type StatefulType struct {
	state int
}

func (st *StatefulType) SetState(ctx context.Context, state int) error {
	st.state = state
	return nil
}

var SetState = Register((*StatefulType).SetState)
