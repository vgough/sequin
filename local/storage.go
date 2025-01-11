package local

import (
	"context"
	"errors"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
)

var ErrOperationAlreadyExists = errors.New("operation already exists")
var ErrOperationNotFound = errors.New("operation not found")
var ErrOperationNotStarted = errors.New("operation not started")
var ErrOperationAlreadyFinished = errors.New("operation already finished")

type OpState struct {
	Checkpoint map[string][]byte
	Results    [][]byte
}

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
		cp *OpState) error

	GetState(ctx context.Context,
		requestID string) (*OpState, error)
}
