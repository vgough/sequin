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

func TestShutdownWorkflow(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	t.Run("Shutdown Running", func(t *testing.T) {
		dropletID := uint64(1000)
		alphaState[dropletID] = true
		virtState[dropletID] = RunningVM

		inc := &ShutdownWorkflow{DropletID: dropletID}
		_, err := autotest.Run(ctx, t, "increment", inc)
		require.NoError(t, err)

		require.False(t, alphaState[dropletID])
		require.Equal(t, ShuttingDownVM, virtState[dropletID])
	})
}

func init() {
	// All workflows must be registered.
	gob.Register(&ShutdownWorkflow{})
}

// ShutdownWorkflow is a mock workflow for shutting down Droplets.
type ShutdownWorkflow struct {
	DropletID uint64
}

// Run implements the Shutdown workflow.
// This is the asynchronous portion of shutdown, which occurs after the request
// has been accepted.  This means we don't need to check for user authorization.
//
// Steps:
//     1. update alpha state
//     2. libvirt: graceful shutdown request
//     4. libvirt: poweroff request if graceful shutdown failed.
//     6. lifecycle event creation
func (wf *ShutdownWorkflow) Run(ctx context.Context, rt workflow.Runtime) {
	var running bool
	rt.Do(func() {
		var err error
		running, err = getAlphaState(ctx, wf.DropletID)
		if err != nil {
			rt.Fail(err)
		}
	}, workflow.State(&running))

	if !running {
		return // Success.  Nothing to shutdown.
	}

	var attemptedVirtShutdown bool
	// We can try to roll back alpha state if we fail before we've asked
	// libvirt to shut down.  After that point, we can't be sure that the
	// droplet will still be running, so we won't try to roll back.
	rt.OnRollback(func() {
		if !attemptedVirtShutdown {
			if err := setAlphaState(ctx, wf.DropletID, false, running); err != nil {
				rt.Fail(err) // rollbacks can fail and be retried!
			}
		}
	})

	// Set new alpha state.
	rt.Do(func() {
		if err := setAlphaState(ctx, wf.DropletID, running, false); err != nil {
			rt.Fail(err)
		}
	}, workflow.Name("shutdown_state"))

	// Persist that we're making libvirt call, so that we don't rollback alpha
	// beyond this point.
	rt.Do(func() {
		attemptedVirtShutdown = true
	}, workflow.State(&attemptedVirtShutdown))

	// Try libvirt shutdown, or fallback to poweroff.
	rt.Do(func() {
		if err := libvirtShutdown(ctx, wf.DropletID); err != nil {
			// shutdown failed, so try falling back to power-off.
			if err := libvirtPowerOff(ctx, wf.DropletID); err != nil {
				rt.Fail(err)
			}
		}
	}, workflow.Name("shutdown_virt"))

	// Finally, send lifecycle event.  This cannot be rolled back.
	// TODO: specify workflow retry option to give essentially unlimited
	// retries, so that we eventually publish, even if it takes a while.
	rt.Do(func() {
		if err := lifecyclePublish(ctx, wf.DropletID); err != nil {
			rt.Fail(err)
		}
	}, workflow.Name("emit_lifecycle"))
}

type VMState int

const (
	RunningVM VMState = iota
	ShuttingDownVM
	PowerdOffVM
)

var errNoSuchDroplet = errors.New("no such droplet")

// Global state representing libvirt's view of droplets.
var virtState = map[uint64]VMState{}

// Global state representing Alpha view of if a droplet is running or not.
var alphaState = map[uint64]bool{}

func getAlphaState(ctx context.Context, dropletID uint64) (bool, error) {
	state, ok := alphaState[dropletID]
	if !ok {
		return false, job.Unrecoverable(errNoSuchDroplet)
	}
	return state, nil
}

func setAlphaState(ctx context.Context, dropletID uint64, oldState, newState bool) error {
	state, ok := alphaState[dropletID]
	if ok && state != oldState {
		return job.Unrecoverable(errors.New("state mismatch"))
	}
	alphaState[dropletID] = newState
	return nil
}

// Mock function to stand in for a libvirt shutdown request.
// First result indicates if call succeeded.  If false, then retry.
func libvirtShutdown(ctx context.Context, dropletID uint64) error {
	state, ok := virtState[dropletID]
	if !ok {
		return job.Unrecoverable(errNoSuchDroplet)
	}

	if state == RunningVM {
		virtState[dropletID] = ShuttingDownVM
	}
	return nil
}

// Mock function to stand in for a libvirt poweroff request.
// First result indicates if call succeeded.  If false, then retry.
func libvirtPowerOff(ctx context.Context, dropletID uint64) error {
	state, ok := virtState[dropletID]
	if !ok {
		return job.Unrecoverable(errors.New("no such droplet"))
	}

	if state != PowerdOffVM {
		virtState[dropletID] = PowerdOffVM
	}
	return nil
}

// Mock function to stand in for publishing a lifecycle event.
func lifecyclePublish(ctx context.Context, dropletID uint64) error {
	return nil
}
