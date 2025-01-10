package storage

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/anypb"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
	"github.com/vgough/sequin/storage/ent"
	"github.com/vgough/sequin/storage/ent/enttest"
	"github.com/vgough/sequin/storage/ent/migrate"
)

func TestObjectStore(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()

	ctx := context.Background()
	os := NewObjectStore(client)
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
