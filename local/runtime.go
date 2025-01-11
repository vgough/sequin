package local

import (
	"context"
	"fmt"
	"reflect"

	"github.com/vgough/sequin/internal"
)

type Runtime struct {
	store     Store
	requestID string
	state     *OpState

	registered map[string]reflect.Value
}

var _ internal.OpRuntime = (*Runtime)(nil)

func NewRuntime(s Store, requestID string, state *OpState) *Runtime {
	return &Runtime{
		store:      s,
		requestID:  requestID,
		state:      state,
		registered: make(map[string]reflect.Value),
	}
}

func (r *Runtime) RegisterState(name string, state any) {
	sv := reflect.ValueOf(state)
	if sv.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("state %s is not a pointer", name))
	}
	if !sv.Elem().CanSet() {
		panic(fmt.Sprintf("state %s is not settable", name))
	}
	r.registered[name] = sv
}

func (r *Runtime) Checkpoint(ctx context.Context) error {
	if r.state.Checkpoint == nil {
		r.state.Checkpoint = make(map[string][]byte)
	}
	for name, state := range r.registered {
		data, err := internal.Encode(state)
		if err != nil {
			return fmt.Errorf("failed to checkpoint state %s: %w", name, err)
		}
		r.state.Checkpoint[name] = data
	}
	return r.store.SetState(ctx, r.requestID, r.state)
}

func (r *Runtime) Restore(ctx context.Context) error {
	for name, data := range r.state.Checkpoint {
		state := r.registered[name]
		v, err := internal.Decode(data, state.Type())
		if err != nil {
			return fmt.Errorf("failed to restore state %s: %w", name, err)
		}
		state.Elem().Set(v.Elem())
	}
	return nil
}
