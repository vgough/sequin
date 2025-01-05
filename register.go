// Package sequin allows registering functions for stateful operation.
package sequin

import (
	"reflect"

	"github.com/vgough/sequin/internal"
	"github.com/vgough/sequin/registry"
)

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

// RegisterOpt is an option for Register.
type RegisterOpt func(*registry.Endpoint) error

// GlobalID marks the function as having a globally unique ID.
// The default is to scope the ID within any enclosing execution.
func GlobalID() RegisterOpt {
	return func(ep *registry.Endpoint) error {
		ep.Metadata[internal.GlobalIDGen] = true
		return nil
	}
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
