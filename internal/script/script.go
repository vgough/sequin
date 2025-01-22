// Package script provides a script interpreter.
package script

import (
	"context"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"

	"github.com/vgough/sequin/internal/script/exports"
)

// New creates a new script interpreter.
func New() *Script {
	i := interp.New(interp.Options{})
	_ = i.Use(stdlib.Symbols)
	_ = i.Use(exports.Symbols)
	return &Script{i: i}
}

// Script is a script interpreter.
type Script struct {
	i *interp.Interpreter
}

// Eval evaluates the given code.
func (s *Script) Eval(ctx context.Context, code string) (any, error) {
	return s.i.EvalWithContext(ctx, code)
}
