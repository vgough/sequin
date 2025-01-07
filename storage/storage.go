package storage

import (
	"context"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
)

type Store interface {
	AddRequest(ctx context.Context,
		req *sequinv1.Operation,
		state *sequinv1.OperationState) error
	SetState(ctx context.Context,
		requestID string,
		state *sequinv1.OperationState) error
	GetState(ctx context.Context,
		requestID string) (*sequinv1.OperationState, error)
	GetRequest(ctx context.Context,
		requestID string) (*sequinv1.Operation, *sequinv1.OperationState, error)

	// TODO: ListRequests.
}
