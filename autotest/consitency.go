package autotest

import (
	"context"

	"github.com/stretchr/testify/assert"
	"github.com/vgough/sequin/workflow"
)

type testConsistency struct {
	twf *testWorkflow
}

func (wf *testConsistency) Run(ctx context.Context, rt workflow.Runtime) {
	// Run through the real workflow a couple times, so that we can
	// check that the resulting state is deterministic.
	wf1 := wf.twf.newInstance(rt)
	wf1.Run(ctx, rt)

	wf2 := wf.twf.newInstance(rt)
	wf2.Run(ctx, rt)

	assert.Equal(wf.twf.t, wf1, wf2,
		"expecting two runs of workflow to produce consistent output state")
}

func (wf *testConsistency) Result() error {
	return nil
}
