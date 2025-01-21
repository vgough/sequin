package local

import (
	"context"
	"errors"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
)

// ErrOperationAlreadyExists is returned when an operation already exists.
var ErrOperationAlreadyExists = errors.New("operation already exists")

// ErrOperationNotFound is returned when an operation is not found.
var ErrOperationNotFound = errors.New("operation not found")

// ErrOperationNotStarted is returned when an operation is not started.
var ErrOperationNotStarted = errors.New("operation not started")

// ErrOperationAlreadyFinished is returned when an operation is already finished.
var ErrOperationAlreadyFinished = errors.New("operation already finished")

// Store is the interface for a storage backend.
type Store interface {
	// AddOperation adds an operation to the store.
	// If the operation already exists, it returns the existing operation.
	AddOperation(ctx context.Context,
		req *sequinv1.Operation) (created bool, err error)

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
