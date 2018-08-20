package workflow

import (
	"context"
	"encoding/gob"
)

// TwoPhaseCommit adds a two-phase commit operation on a resource.  This
// consists of a non-idempotent reserve step which returns a resource id, and
// idempotent commit and release functions.
//
// Note that it is possible for resources to be reserved but not committed or
// released, so if you can't afford to loose resources in case of system
// failure, then there should be some out-of-band garbage collection for
// reserved but uncommitted resources to reclaim unused resources.
func TwoPhaseCommit(rt Runtime, reserve func() (id string),
	commit, release func(id string), opts ...Opt) {
	rt.Embed(&twoPhaseCommitWF{
		reserve: reserve,
		commit:  commit,
		release: release,
	}, opts...)
}

func init() {
	gob.Register(&twoPhaseCommitWF{})
}

// twoPhaseCommitWF is a workflow that implements two-phase commit on a resource.
type twoPhaseCommitWF struct {
	ID string

	reserve         func() (id string)
	commit, release func(id string)
}

// Run implements Workflow.
func (wf *twoPhaseCommitWF) Run(ctx context.Context, rt Runtime) {
	// AtLeastOnce commits state before proceeding.
	rt.Do(func() { wf.ID = wf.reserve() }, Name("reserve"))

	// Because the rollback is stateful, it might not be executed in edge cases
	// even if reserve succeeded.  Hence, two-phase commit should be pairted
	// with garbage collection of reserved but uncommitted resources.
	rt.OnRollback(func() { wf.release(wf.ID) })

	rt.Do(func() { wf.commit(wf.ID) }, Name("commit"))
}
