package local

import (
	"context"
	"sync"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
)

type MemStore struct {
	mu  sync.Mutex
	ops map[string]*opInfo
}

var _ Store = &MemStore{}

type opInfo struct {
	def    *sequinv1.Operation
	state  *sequinv1.OperationState
	result *sequinv1.OperationResult
}

func NewMemStore() *MemStore {
	return &MemStore{
		ops: make(map[string]*opInfo),
	}
}

func (s *MemStore) AddOperation(ctx context.Context,
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

func (s *MemStore) GetOperation(ctx context.Context,
	requestID string) (*sequinv1.Operation, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	op, ok := s.ops[requestID]
	if !ok {
		return nil, ErrOperationNotFound
	}
	return op.def, nil
}

func (s *MemStore) SetState(ctx context.Context,
	requestID string,
	state *sequinv1.OperationState) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ops[requestID].state = state
	return nil
}

func (s *MemStore) GetState(ctx context.Context,
	requestID string) (*sequinv1.OperationState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	op, ok := s.ops[requestID]
	if !ok {
		return nil, ErrOperationNotFound
	}
	return op.state, nil
}

func (s *MemStore) SetResult(ctx context.Context,
	requestID string,
	result *sequinv1.OperationResult) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.ops[requestID].result = result
	return nil
}

func (s *MemStore) GetResult(ctx context.Context,
	requestID string) (bool, *sequinv1.OperationResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	op, ok := s.ops[requestID]
	if !ok {
		return false, nil, ErrOperationNotFound
	}
	return true, op.result, nil
}
