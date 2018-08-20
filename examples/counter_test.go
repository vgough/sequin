package examples

import (
	"context"
	"encoding/gob"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/vgough/sequin/workflow"
	"github.com/vgough/sequin/autotest"
)

func TestIncrementWorkflow(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	inc := &IncrementWorkflow{Delta: 1}
	_, err := autotest.Run(ctx, t, "counter_test", inc)
	require.NoError(t, err)
}

func init() {
	// All workflows must be registered.
	gob.Register(&IncrementWorkflow{})
}

// IncrementWorkflow is a workflow example.
type IncrementWorkflow struct {
	Delta int64
}

func (wf *IncrementWorkflow) Run(ctx context.Context, run workflow.Runtime) {
	// Stage which gets the value.
	var old int64
	// Persist option ensures that local variable side effects are persisted.
	run.Do(func() { old = getGlobal() },
		workflow.State(&old), workflow.Name("get_value"))

	// Not necessary to put everything inside of an Apply operation.
	// The difference is that this code will be run on every retry, where
	// the apply operations will use cached values.  There's also no undo.
	// This is a good place for `connectivity` code, which connects outputs
	// of one operation to inputs of the next operations.
	newVal := old + wf.Delta

	// Increment list of who changed the value, idempotently.
	// Both apply and release functions can be executed any number of times
	// and the result is the same.
	var id uuid.UUID
	run.Do(func() {
		var err error
		id, err = uuid.NewDCEPerson()
		if err != nil {
			run.Fail(err)
		}
		changers.Store(id, true)
	}, workflow.State(&id), workflow.Name("inc_count"))
	run.OnRollback(func() { changers.Delete(id) })

	// Stage which sleeps for a randomized period of time.
	// In an apply function so that this isn't rerun on a retry.
	run.Do(func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}, workflow.Name("sleep"))

	// Stage which sets a new value, or else updates a numConflicts metric.
	// Always succeeds or fails workflow, so there's no revert operation.
	run.Do(func() {
		if !compareAndSet(old, newVal) {
			// Restart entire workflow since our optimistic lock failed.
			run.Fail(errors.Errorf("optimistic locking failed (%v -> %v)", old, newVal))
		}
	}, workflow.RollbackOnFailure(), workflow.Name("save"))
}

var globalValue int64
var changers sync.Map

func compareAndSet(old, v int64) bool {
	return atomic.CompareAndSwapInt64(&globalValue, old, v)
}

func getGlobal() int64 {
	return globalValue
}
