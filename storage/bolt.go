package storage

import (
	"context"
	"fmt"

	"github.com/timshannon/bolthold"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
)

type BoltStore struct {
	store *bolthold.Store
}

var _ Store = &BoltStore{}

type Operation struct {
	ID        string
	Done      bool
	Operation []byte
	State     []byte
}

func NewBoltStore(filename string) (*BoltStore, error) {
	// TODO: replace codec with proto support?
	opts := &bolthold.Options{
		Encoder: bolthold.DefaultEncode,
		Decoder: bolthold.DefaultDecode,
	}
	store, err := bolthold.Open(filename, 0666, opts)
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return &BoltStore{store: store}, nil
}

// AddRequest implements Store.
func (b *BoltStore) AddRequest(ctx context.Context, req *sequinv1.Operation,
	state *sequinv1.OperationState) error {

	b.store.
		panic("unimplemented")
}

// GetRequest implements Store.
func (b *BoltStore) GetRequest(ctx context.Context,
	requestID string) (*sequinv1.Operation, *sequinv1.OperationState, error) {
	panic("unimplemented")
}

// GetState implements Store.
func (b *BoltStore) GetState(ctx context.Context,
	requestID string) (*sequinv1.OperationState, error) {
	panic("unimplemented")
}

// SetState implements Store.
func (b *BoltStore) SetState(ctx context.Context, requestID string,
	state *sequinv1.OperationState) error {
	panic("unimplemented")
}
