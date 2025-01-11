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
	def        *sequinv1.Operation
	checkpoint map[string][]byte
	result     [][]byte
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
		return false, nil
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
	state *OpState) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if state.Checkpoint != nil {
		s.ops[requestID].checkpoint = state.Checkpoint
	}
	if state.Results != nil {
		s.ops[requestID].result = state.Results
	}
	return nil
}

func (s *MemStore) GetState(ctx context.Context,
	requestID string) (*OpState, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	op, ok := s.ops[requestID]
	if !ok {
		return nil, ErrOperationNotFound
	}
	return &OpState{
		Checkpoint: op.checkpoint,
		Results:    op.result,
	}, nil
}
