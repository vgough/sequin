package storage

import (
	"context"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
)

type Store interface {
	AddOperation(ctx context.Context,
		req *sequinv1.Operation,
		submitter string) error

	GetOperation(ctx context.Context,
		requestID string) (*sequinv1.Operation, error)

	// TODO: ListOperations.

	SetState(ctx context.Context,
		requestID string,
		state *sequinv1.OperationState) error

	GetState(ctx context.Context,
		requestID string) (*sequinv1.OperationState, error)

	SetResult(ctx context.Context,
		requestID string,
		result *sequinv1.OperationResult) error

	GetResult(ctx context.Context,
		requestID string) (bool, *sequinv1.OperationResult, error)
}
