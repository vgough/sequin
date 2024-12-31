package internal

import "context"

// MDKey is a context metadata helper.
type MDKey[T any] struct{}

// Get returns the metadata value stored in the context, or else the zero value.
func (md MDKey[T]) Get(ctx context.Context) T {
	v := ctx.Value(md)
	if v == nil {
		var zero T
		return zero
	}
	return v.(T)
}

// Set stores the metadata value in the context.
func (md MDKey[T]) Set(ctx context.Context, v T) context.Context {
	return context.WithValue(ctx, md, v)
}
