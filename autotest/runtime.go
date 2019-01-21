package autotest

import (
	"container/list"
	"context"

	"github.com/vgough/sequin/workflow"
)

// OpFN is an alias to make the signatures look the same as in the workflow
// package.
type OpFN = workflow.OpFN

// Opt is an alias to make the signatures look the same as in the workflow
// package.
type Opt = workflow.Opt

type testRuntime struct {
	t         Testing
	name      string
	base      workflow.Runtime
	nextSteps *list.List
}

var _ workflow.Runtime = &testRuntime{}

func (run *testRuntime) ID() string {
	return run.base.ID()
}

func (run *testRuntime) Context() context.Context {
	return run.base.Context()
}

// Do implements Runtime.
func (run *testRuntime) Do(apply OpFN, opts ...Opt) {
	run.base.Do(apply, opts...)
}

// Embed implements Runtime.
func (run *testRuntime) Embed(wf workflow.Workflow, opts ...Opt) {
	// fqn := path.Join(run.name, name)
	run.base.Embed(wf, opts...)
}

func (run *testRuntime) OnRollback(fn OpFN) {
	run.base.OnRollback(fn)
}

// Fail implements Runtime.
func (run *testRuntime) Fail(err error) {
	run.base.Fail(err)
}
