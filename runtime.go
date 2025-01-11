package sequin

import (
	"context"
	"reflect"

	"github.com/vgough/sequin/internal"
	"github.com/vgough/sequin/registry"
)

var runtimeMD internal.MDKey[Runtime]

// Runtime is the interface for a runtime that can execute operations.
type Runtime interface {
	Exec(ep *registry.Endpoint, args []reflect.Value) []reflect.Value
}

// WithRuntime stores the runtime in the context.
func WithRuntime(ctx context.Context, rt Runtime) context.Context {
	return runtimeMD.Set(ctx, rt)
}

// GetRuntime retrieves the runtime from the context.
func GetRuntime(ctx context.Context) Runtime {
	return runtimeMD.Get(ctx)
}

type OpRuntimeOpt func(rt internal.OpRuntime) error

func OpInit(ctx context.Context, opts ...OpRuntimeOpt) error {
	rt := internal.GetOpRuntime(ctx)
	if rt == nil {
		rt = internal.NoOpRuntime{}
	}
	for _, opt := range opts {
		if err := opt(rt); err != nil {
			return err
		}
	}
	return rt.Restore(ctx)
}

func Persist[T any](name string, state T) OpRuntimeOpt {
	return func(rt internal.OpRuntime) error {
		rt.RegisterState(name, state)
		return nil
	}
}

// Checkpoint stores the current state of any registered variables.
func Checkpoint(ctx context.Context) error {
	rt := internal.GetOpRuntime(ctx)
	if rt == nil {
		return nil
	}
	return rt.Checkpoint(ctx)
}
