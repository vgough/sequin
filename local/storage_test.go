package local

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
)

func TestObjectStore(t *testing.T) {
	ctx := context.Background()
	os := NewMemStore()
	created, err := os.AddOperation(ctx, &sequinv1.Operation{
		RequestId: "123",
		Detail: &anypb.Any{
			TypeUrl: "test",
			Value:   []byte("test"),
		},
	})
	require.NoError(t, err)
	require.True(t, created)
	err = os.SetState(ctx, "123", &OpState{
		Checkpoint: map[string][]byte{
			"test": []byte("test"),
		},
	})
	require.NoError(t, err)

	cp, err := os.GetState(ctx, "123")
	require.NoError(t, err)
	require.Equal(t, cp.Checkpoint, map[string][]byte{
		"test": []byte("test"),
	})
}
