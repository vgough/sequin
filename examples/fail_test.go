package examples

import (
	"context"
	"encoding/gob"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/vgough/sequin/workflow"
	"github.com/vgough/sequin/autotest"
	"github.com/vgough/sequin/job"
)

func TestFailWorkflow(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wf := &failWorkflow{}
	out, err := autotest.Run(ctx, t, "fail_test", wf)
	require.Error(t, err)
	require.Equal(t, "aborted", out.(*failWorkflow).State)
}

func init() {
	// All workflows must be registered.
	gob.Register(&failWorkflow{})
}

// failWorkflow is a workflow example.
type failWorkflow struct {
	State string
}

func (wf *failWorkflow) Run(ctx context.Context, run workflow.Runtime) {
	run.OnRollback(func() {
		run.Log().Warn("in rollback")
		wf.State = "aborted"
	})

	run.Do(func() { wf.State = "step1" })
	if wf.State != "step1" {
		panic("invalid workflow state")
	}

	// Run a step which fails, forcing a rollback.
	run.Do(func() {
		run.Fail(job.Unrecoverable(errors.New("must fail")))
	})
}
