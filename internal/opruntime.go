package internal

import "context"

var opRuntimeMD MDKey[OpRuntime]

type OpRuntime interface {
	Checkpoint(ctx context.Context) error
	RegisterState(name string, state any)
	Restore(ctx context.Context) error
}

func WithOpRuntime(ctx context.Context, rt OpRuntime) context.Context {
	return opRuntimeMD.Set(ctx, rt)
}

func GetOpRuntime(ctx context.Context) OpRuntime {
	return opRuntimeMD.Get(ctx)
}

type NoOpRuntime struct{}

func (NoOpRuntime) RegisterState(name string, state any) {
}

func (NoOpRuntime) Checkpoint(ctx context.Context) error {
	return nil
}

func (NoOpRuntime) Restore(ctx context.Context) error {
	return nil
}
