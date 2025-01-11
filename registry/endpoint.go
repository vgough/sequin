package registry

import (
	"context"
	"errors"
	"reflect"
	"runtime"
)

var ContextType = reflect.TypeFor[context.Context]()
var ErrorType = reflect.TypeFor[error]()

var ErrEndpointNotFound = errors.New("endpoint not found")

type FeatureFlag int

const (
	// FeatureCheckpoint enables checkpointing for the operation.
	FeatureCheckpoint FeatureFlag = 1 << iota
	// FeatureBackground enables automatic background execution for the operation.
	FeatureBackground
)

// Endpoint stores information about a registered function.
type Endpoint struct {
	Name          string
	LocalFN       reflect.Value  // Local function to call.
	ContextIndex  int            // Index of the context.Context argument.
	InputTypes    []reflect.Type // Argument types of the function.
	OutputTypes   []reflect.Type // Return types of the function.
	MethodBinding reflect.Value  // Fixed method to bind to calls.
	FeatureFlags  FeatureFlag    // Feature flags.

	Metadata map[string]interface{} // Arbitrary metadata.

	// TODO: optional configuration overrides when used as an embedded component
	// rather than a top-level function.
	Embedded *EmbeddedOverrides
}

func (ep *Endpoint) SetMetadata(key string, value interface{}) {
	if ep.Metadata == nil {
		ep.Metadata = make(map[string]interface{})
	}
	ep.Metadata[key] = value
}

func (ep *Endpoint) HasFeature(flag FeatureFlag) bool {
	return ep.FeatureFlags&flag != 0
}

func (ep *Endpoint) EnableFeature(flag FeatureFlag) {
	ep.FeatureFlags |= flag
}

type EmbeddedOverrides struct {
	FeatureFlags FeatureFlag
	Metadata     map[string]interface{}
}

func (eo *EmbeddedOverrides) EnableFeature(flag FeatureFlag) {
	eo.FeatureFlags |= flag
}

func (eo *EmbeddedOverrides) SetMetadata(key string, value interface{}) {
	if eo.Metadata == nil {
		eo.Metadata = make(map[string]interface{})
	}
	eo.Metadata[key] = value
}

// NewEndpoint creates a new endpoint from a function.
//
// The function must take a context.Context as one of the arguments, and
// must return an error as the last result.
//
// If the function does not have a valid signature, an error is returned.
func NewEndpoint(fn reflect.Value) (*Endpoint, error) {
	fnT := fn.Type()
	switch {
	case fnT.Kind() != reflect.Func:
		return nil, errors.New("sequin.Register requires a function argument")
	case fnT.NumIn() < 1:
		return nil, errors.New("function requires at least one argument")
	case fnT.NumOut() < 1:
		return nil, errors.New("function must return at least one result")
	case fnT.Out(fnT.NumOut()-1) != ErrorType:
		return nil, errors.New("function must return an error as the last result")
	}

	fnInfo := runtime.FuncForPC(fn.Pointer())
	ep := &Endpoint{
		Name:         fnInfo.Name(),
		LocalFN:      fn,
		ContextIndex: -1,
		InputTypes:   make([]reflect.Type, fnT.NumIn()),
		OutputTypes:  make([]reflect.Type, fnT.NumOut()),
		Metadata:     make(map[string]interface{}),
	}
	for i := range fnT.NumIn() {
		ep.InputTypes[i] = fnT.In(i)
	}
	for i := range fnT.NumOut() {
		ep.OutputTypes[i] = fnT.Out(i)
	}

	// Determine how to get the context.
	// Anything before the context must be injected, while the remaining
	// arguments are passed.
	// This allows for method functions to be registered.
	for i := 0; i < fnT.NumIn(); i++ {
		if fnT.In(i) == ContextType {
			ep.ContextIndex = i
			break
		}
	}
	return ep, nil
}

// Register adds an endpoint to the global registry.
func Register(ep *Endpoint) error {
	if ep.Name == "" {
		return errors.New("endpoint name cannot be empty")
	}
	if _, ok := registry[ep.Name]; ok {
		return errors.New("endpoint already registered")
	}

	registry[ep.Name] = ep
	return nil
}

// RegisteredEndpoints returns a list of names of all registered endpoints.
func RegisteredEndpoints() []string {
	endpoints := make([]string, 0, len(registry))
	for name := range registry {
		endpoints = append(endpoints, name)
	}
	return endpoints
}

// GetEndpoint returns the endpoint with the given name.
// Returns nil if the endpoint is not found.
func GetEndpoint(name string) *Endpoint {
	return registry[name]
}

var registry = map[string]*Endpoint{}

func (ep *Endpoint) GetContext(args []reflect.Value) context.Context {
	return args[ep.ContextIndex].Interface().(context.Context)
}

func (ep *Endpoint) SetContext(ctx context.Context, args []reflect.Value) {
	args[ep.ContextIndex] = reflect.ValueOf(ctx)
}

// MakeError returns a slice of reflect.Value with the error value set.
func (ep *Endpoint) MakeError(err error) []reflect.Value {
	return MakeError(err, ep.OutputTypes)
}

// Exec calls the local function with the given context and arguments.
func (ep *Endpoint) Exec(args []reflect.Value) []reflect.Value {
	return ep.LocalFN.Call(args)
}

func (ep *Endpoint) GetError(results []reflect.Value) error {
	if len(results) == 0 {
		return nil
	}
	last := results[len(results)-1]
	if last.IsNil() {
		return nil
	}
	return last.Interface().(error)
}

// MakeError returns a slice of reflect.Value with the error value set.
func MakeError(err error, outputTypes []reflect.Type) []reflect.Value {
	out := make([]reflect.Value, len(outputTypes))
	for i, ot := range outputTypes {
		if ot == ErrorType {
			out[i] = reflect.ValueOf(err)
		} else {
			out[i] = reflect.Zero(ot)
		}
	}
	return out
}
