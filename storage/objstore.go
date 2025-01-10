package storage

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"log"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
	"github.com/vgough/sequin/storage/ent"
	"github.com/vgough/sequin/storage/ent/operation"

	"google.golang.org/protobuf/proto"
)

type ObjStore struct {
	db        *ent.Client
	numShards int64
}

type ObjectStoreOption func(*ObjStore)

func WithNumShards(numShards int64) ObjectStoreOption {
	return func(o *ObjStore) {
		o.numShards = numShards
	}
}

// NewTestObjectStore creates a new in-memory object store for testing.
// If using this, you must import the sqlite driver:
//
//	import _ "github.com/mattn/go-sqlite3"
func NewTestObjectStore(opts ...ObjectStoreOption) *ObjStore {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return NewObjectStore(client, opts...)
}

func NewObjectStore(db *ent.Client, opts ...ObjectStoreOption) *ObjStore {
	o := &ObjStore{db: db, numShards: 1}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

var _ Store = (*ObjStore)(nil)

// AddOperation implements Store.
func (o *ObjStore) AddOperation(ctx context.Context,
	req *sequinv1.Operation) (bool, error) {

	detail, err := proto.Marshal(req)
	if err != nil {
		return false, fmt.Errorf("marshal failed: %w", err)
	}

	shard := o.shardFromRequestID(req.RequestId)

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return false, fmt.Errorf("begin tx failed: %w", err)
	}
	defer tx.Rollback()

	entop, err := tx.Operation.Query().
		Select(operation.FieldDetail).
		Where(operation.Shard(shard), operation.RequestID(req.RequestId)).
		Only(ctx)

	switch {
	case err == nil:
		detail := &sequinv1.Operation{}
		err = proto.Unmarshal(entop.Detail, detail)
		if err != nil {
			return false, fmt.Errorf("unmarshal failed: %w", err)
		}
		return false, nil

	case ent.IsNotFound(err):
		// Operation doesn't exist, so we can add it.
	default:
		return false, fmt.Errorf("get failed: %w", err)
	}

	_, err = tx.Operation.Create().
		SetRequestID(req.RequestId).
		SetShard(shard).
		SetDetail(detail).
		Save(ctx)

	if err != nil {
		return false, fmt.Errorf("add failed: %w", err)
	}

	return true, tx.Commit()
}

// func (o *ObjStore) AddLabel(ctx context.Context, requestID string, name, value string) error {
// 	entop, err := o.db.Operation.Get(ctx, requestID)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = o.db.Label.Create().
// 		SetName(name).
// 		SetValue(value).
// 		AddOperation(entop).
// 		Save(ctx)
// 	return err
// }

func (o *ObjStore) shardFromRequestID(requestID string) int64 {
	hash := fnv.New64()
	hash.Write([]byte(requestID))
	// Set bottom two bits to 0 to allow for shard changes in the future.
	id := hash.Sum64() &^ 3
	return int64(id) % o.numShards
}

// GetOperation implements Store.
func (o *ObjStore) GetOperation(ctx context.Context,
	requestID string) (*sequinv1.Operation, error) {

	shard := o.shardFromRequestID(requestID)

	entop, err := o.db.Operation.Query().
		Select(operation.FieldDetail).
		Where(operation.Shard(shard), operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrOperationNotFound
		}
		return nil, fmt.Errorf("get failed: %w", err)
	}

	detail := &sequinv1.Operation{}
	err = proto.Unmarshal(entop.Detail, detail)
	if err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return detail, nil
}

// GetState implements Store.
func (o *ObjStore) GetState(ctx context.Context,
	requestID string) (*sequinv1.OperationState, error) {

	shard := o.shardFromRequestID(requestID)

	entop, err := o.db.Operation.Query().
		Select(operation.FieldState).
		Where(operation.Shard(shard), operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrOperationNotFound
		}
		return nil, fmt.Errorf("get failed: %w", err)
	}

	state := &sequinv1.OperationState{}
	err = proto.Unmarshal(entop.State, state)
	if err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return state, nil
}

// SetState implements Store.
func (o *ObjStore) SetState(ctx context.Context,
	requestID string, state *sequinv1.OperationState) error {

	shard := o.shardFromRequestID(requestID)

	stateData, err := proto.Marshal(state)
	if err != nil {
		return fmt.Errorf("marshal failed: %w", err)
	}

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx failed: %w", err)
	}
	defer tx.Rollback()

	// Check that the operation exists and is not done.
	entop, err := tx.Operation.Query().
		Select(operation.FieldResult, operation.FieldState).
		Where(operation.Shard(shard), operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("get failed: %w", err)
	}
	if entop.Result != nil {
		return fmt.Errorf("operation already done")
	}

	if entop.State != nil {
		var oldState sequinv1.OperationState
		if err := proto.Unmarshal(entop.State, &oldState); err != nil {
			return fmt.Errorf("unmarshal failed: %w", err)
		}
		if oldState.UpdateId >= state.UpdateId {
			return errors.New("update ID is not increasing")
		}
	}

	_, err = entop.Update().
		SetState(stateData).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	return tx.Commit()
}

func (o *ObjStore) SetResult(ctx context.Context,
	requestID string, result *sequinv1.OperationResult) error {

	resultData, err := proto.Marshal(result)
	if err != nil {
		return fmt.Errorf("marshal failed: %w", err)
	}

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx failed: %w", err)
	}
	defer tx.Rollback()

	entop, err := tx.Operation.Query().
		Select(operation.FieldResult).
		Where(operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrOperationNotFound
		}
		return fmt.Errorf("get failed: %w", err)
	}
	if entop.Result != nil {
		return ErrOperationAlreadyFinished
	}

	_, err = entop.Update().
		SetResult(resultData).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	return tx.Commit()
}

// GetResult implements Store.
func (o *ObjStore) GetResult(ctx context.Context,
	requestID string) (bool, *sequinv1.OperationResult, error) {

	entop, err := o.db.Operation.Query().
		Select(operation.FieldResult).
		Where(operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil, ErrOperationNotFound
		}
		return false, nil, fmt.Errorf("get failed: %w", err)
	}

	if entop.Result == nil {
		return false, nil, nil
	}
	result := &sequinv1.OperationResult{}
	err = proto.Unmarshal(entop.Result, result)
	if err != nil {
		return false, nil, fmt.Errorf("unmarshal failed: %w", err)
	}
	return true, result, nil
}
