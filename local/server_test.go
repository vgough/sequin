package local

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vgough/sequin"
)

func TestServer(t *testing.T) {
	s := NewServer()
	ctx := sequin.WithRuntime(context.Background(), s)

	ok, err := IsEven(ctx, 0)
	require.NoError(t, err)
	require.True(t, ok)
}

var IsEven = sequin.Register(isEven)

func isEven(_ context.Context, in int) (bool, error) {
	if in < 0 {
		return false, errors.New("negative values not supported")
	}
	if in%2 == 0 {
		return true, nil
	}
	return false, nil
}
