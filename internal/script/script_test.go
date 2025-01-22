package script

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {
	s := New()
	_, err := s.Eval(context.Background(), `
		package main

		func Hello() string {
			return "Hello, World!"
		}
	`)
	require.NoError(t, err)

	hello, err := s.i.Eval("main.Hello()")
	require.NoError(t, err)
	require.Equal(t, "Hello, World!", hello.String())
}
