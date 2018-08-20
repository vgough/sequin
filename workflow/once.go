package workflow

import (
	"context"
	"encoding/gob"

	"github.com/pkg/errors"
	"github.com/vgough/sequin/job"
)

// Once adds a workflow step that implements at-most-once restrictions.
// This is a lot like a two phase commit, with the addition that it fails
// if the first step is attempted more than once.
func Once(rt Runtime, fn OpFN, opts ...Opt) {
	rt.Embed(&atMostOnceWF{fn: fn}, opts...)
}

type atMostOnceWF struct {
	Attempted bool
	fn        OpFN
}

func init() {
	gob.Register(&atMostOnceWF{})
}

func (wf *atMostOnceWF) Run(ctx context.Context, rt Runtime) {
	// Fail workflow if there's any sign that we've attempted it before.
	if wf.Attempted {
		rt.Fail(job.Unrecoverable(errors.New("cannot retry AtMostOnce")))
	}
	// Once this stage commits, we'll no longer be able to restart the workflow.
	rt.Do(func() { wf.Attempted = true }, Name("limit"))

	// Prevent retry from working beyond this point -- OnRollback failure
	// coupled with RollbackOnFailure flag ensures that we won't retry the
	// operation.
	rt.OnRollback(func() {
		rt.Fail(job.Unrecoverable(errors.New("cannot retry AtMostOnce")))
	})
	rt.Do(wf.fn, RollbackOnFailure(), Name("apply"))
}
