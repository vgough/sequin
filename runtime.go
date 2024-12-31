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

func GetRuntime(ctx context.Context) Runtime {
	return runtimeMD.Get(ctx)
}
