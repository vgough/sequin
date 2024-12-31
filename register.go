package sequin

import (
	"reflect"

	"github.com/vgough/sequin/registry"
)

const IDGenMDKey = "sequin.IDGen"

type RegisterOpt func(*registry.Endpoint) error

// WithIDGen sets the function to generate an ID for the runtime.
func WithIDGen(fn func([]reflect.Value) string) RegisterOpt {
	return func(ep *registry.Endpoint) error {
		ep.Metadata[IDGenMDKey] = fn
		return nil
	}
}

// Register a function for stateful operation.
// Returns a wrapper function that can be called to execute the operation.
// The wrapper will look for the runtime in the context to use for execution.
//
// The function must take a context.Context as one of the arguments, and
// must return an error as the last result.
//
// The error return requirement is to allow for error reporting by the runtime.
func Register[T any](fn T, opts ...RegisterOpt) T {
	fnV := reflect.ValueOf(fn)
	ep, err := registry.NewEndpoint(fnV)
	if err != nil {
		panic("sequin.Register called with invalid function: " + err.Error())
	}

	if ep.ContextIndex < 0 {
		panic("function must take a context.Context as one of the arguments")
	}
	for _, opt := range opts {
		if err := opt(ep); err != nil {
			panic(err)
		}
	}

	if err := registry.Register(ep); err != nil {
		panic(err)
	}

	// return runtime wrapper.
	fnT := reflect.TypeOf(fn)
	wrapperFN := reflect.MakeFunc(fnT, contextDispatch(ep))
	var out T
	reflect.ValueOf(&out).Elem().Set(wrapperFN)
	return out
}

func contextDispatch(ep *registry.Endpoint) func([]reflect.Value) []reflect.Value {
	return func(args []reflect.Value) []reflect.Value {
		ctx := ep.GetContext(args)
		rt := GetRuntime(ctx)
		if rt == nil {
			// Whe no runtime is available, call the function directly.
			return ep.Exec(args)
		}
		return rt.Exec(ep, args)
	}
}
