// Package local provides a local runtime for running operations.
package local

import (
	"context"
	"sync"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
)

// MemStore is an in-memory store for operations.
type MemStore struct {
	mu  sync.Mutex
	ops map[string]*opInfo
}

var _ Store = &MemStore{}

// opInfo is the internal representation of an operation.
type opInfo struct {
	def    *sequinv1.Operation
	state  *sequinv1.OperationState
	result *sequinv1.OperationResult
}

// NewMemStore creates a new in-memory store for operations.
func NewMemStore() *MemStore {
	return &MemStore{
		ops: make(map[string]*opInfo),
	}
}

// AddOperation adds a new operation to the store.
func (s *MemStore) AddOperation(_ context.Context,
	req *sequinv1.Operation) (created bool, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.ops[req.RequestId]; ok {
		return false, ErrOperationAlreadyExists
	}

	s.ops[req.RequestId] = &opInfo{
		def: req,
	}
	return true, nil
}

// GetOperation retrieves an operation from the store.
func (s *MemStore) GetOperation(_ context.Context,
	requestID string) (*sequinv1.Operation, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	op, ok := s.ops[requestID]
	if !ok {
		return nil, ErrOperationNotFound
	}
	return op.def, nil
}

// SetState updates the state of an operation.
func (s *MemStore) SetState(_ context.Context,
	requestID string,
	state *sequinv1.OperationState) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ops[requestID].state = state
	return nil
}

// GetState retrieves the state of an operation.
func (s *MemStore) GetState(_ context.Context,
	requestID string) (*sequinv1.OperationState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	op, ok := s.ops[requestID]
	if !ok {
		return nil, ErrOperationNotFound
	}
	return op.state, nil
}

// SetResult updates the result of an operation.
func (s *MemStore) SetResult(_ context.Context,
	requestID string,
	result *sequinv1.OperationResult) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ops[requestID].result = result
	return nil
}

// GetResult retrieves the result of an operation.
func (s *MemStore) GetResult(_ context.Context,
	requestID string) (bool, *sequinv1.OperationResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	op, ok := s.ops[requestID]
	if !ok {
		return false, nil, ErrOperationNotFound
	}
	return true, op.result, nil
}
