package storage

import (
	"context"
	"errors"
	"hash/fnv"
	"log"

	sequinv1 "github.com/vgough/sequin/gen/sequin/v1"
	"github.com/vgough/sequin/storage/ent"
	"github.com/vgough/sequin/storage/ent/operation"

	"google.golang.org/protobuf/proto"
)

type ObjStore struct {
	db *ent.Client
}

// NewTestObjectStore creates a new in-memory object store for testing.
// If using this, you must import the sqlite driver:
//
//	import _ "github.com/mattn/go-sqlite3"
func NewTestObjectStore() *ObjStore {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &ObjStore{db: client}
}

func NewObjectStore(db *ent.Client) *ObjStore {
	return &ObjStore{db: db}
}

var _ Store = (*ObjStore)(nil)

// AddOperation implements Store.
func (o *ObjStore) AddOperation(ctx context.Context, req *sequinv1.Operation,
	submitter string) error {

	detail, err := proto.Marshal(req)
	if err != nil {
		return err
	}

	shard := shardFromRequestID(req.RequestId, 1000)

	_, err = o.db.Operation.Create().
		SetRequestID(req.RequestId).
		SetShard(shard).
		SetDetail(detail).
		SetSubmitter(submitter).
		Save(ctx)
	return err
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

func shardFromRequestID(requestID string, maxShards int64) int64 {
	hash := fnv.New64()
	hash.Write([]byte(requestID))
	// Set bottom two bits to 0 to allow for shard changes in the future.
	id := hash.Sum64() &^ 3
	return int64(id) % maxShards
}

// GetOperation implements Store.
func (o *ObjStore) GetOperation(ctx context.Context,
	requestID string) (*sequinv1.Operation, error) {

	shard := shardFromRequestID(requestID, 1000)

	entop, err := o.db.Operation.Query().
		Select(operation.FieldDetail).
		Where(operation.Shard(shard), operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	detail := &sequinv1.Operation{}
	err = proto.Unmarshal(entop.Detail, detail)
	if err != nil {
		return nil, err
	}

	return detail, nil
}

// GetState implements Store.
func (o *ObjStore) GetState(ctx context.Context,
	requestID string) (*sequinv1.OperationState, error) {

	shard := shardFromRequestID(requestID, 1000)

	entop, err := o.db.Operation.Query().
		Select(operation.FieldState).
		Where(operation.Shard(shard), operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	state := &sequinv1.OperationState{}
	err = proto.Unmarshal(entop.State, state)
	return state, err
}

// SetState implements Store.
func (o *ObjStore) SetState(ctx context.Context,
	requestID string, state *sequinv1.OperationState) error {

	shard := shardFromRequestID(requestID, 1000)

	stateData, err := proto.Marshal(state)
	if err != nil {
		return err
	}

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check that the operation exists and is not done.
	entop, err := tx.Operation.Query().
		Select(operation.FieldResult, operation.FieldState).
		Where(operation.Shard(shard), operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		return err
	}
	if entop.Result != nil {
		return errors.New("operation already done")
	}

	if entop.State != nil {
		var oldState sequinv1.OperationState
		if err := proto.Unmarshal(entop.State, &oldState); err != nil {
			return err
		}
		if oldState.UpdateId >= state.UpdateId {
			return errors.New("update ID is not increasing")
		}
	}

	_, err = entop.Update().
		SetState(stateData).
		Save(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (o *ObjStore) SetResult(ctx context.Context,
	requestID string, result *sequinv1.OperationResult) error {

	resultData, err := proto.Marshal(result)
	if err != nil {
		return err
	}

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	entop, err := tx.Operation.Query().
		Select(operation.FieldResult).
		Where(operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		return err
	}
	if entop.Result != nil {
		return errors.New("operation already done")
	}

	_, err = entop.Update().
		SetResult(resultData).
		Save(ctx)
	return err
}

// GetResult implements Store.
func (o *ObjStore) GetResult(ctx context.Context,
	requestID string) (bool, *sequinv1.OperationResult, error) {

	entop, err := o.db.Operation.Query().
		Select(operation.FieldResult).
		Where(operation.RequestID(requestID)).
		Only(ctx)
	if err != nil {
		return false, nil, err
	}

	if entop.Result == nil {
		return false, nil, nil
	}
	result := &sequinv1.OperationResult{}
	err = proto.Unmarshal(entop.Result, result)
	return true, result, err
}
