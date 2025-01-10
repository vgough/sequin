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
	err = os.SetState(ctx, "123", &sequinv1.OperationState{
		UpdateId: 1,
	})
	require.NoError(t, err)

	state, err := os.GetState(ctx, "123")
	require.NoError(t, err)
	require.Equal(t, state.UpdateId, int64(1))
}
