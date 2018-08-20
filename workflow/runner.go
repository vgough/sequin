package workflow

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"path"

	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/vgough/sequin/job"
	"go.opencensus.io/trace"
)

func init() {
	gob.Register(&JobWrapper{})
}

// JobWrapper implements job.Runnable, allowing workflows to be scheduled
// and executed with the job scheduler.
type JobWrapper struct {
	Name string

	Workflow Workflow

	// Root is the top-level execution context.
	// It is expanded with child nodes as execution progresses.
	Root *ExecContext

	AttemptCount int
}

// Run implements job.Runnable, which in turn runs the wrapped workflow.
func (j *JobWrapper) Run(ctx context.Context, jr job.Runtime) error {
	ctx, span := trace.StartSpan(ctx, j.Name)

	ec := j.Root
	if ec == nil {
		ec = &ExecContext{
			Name: j.Name,
			wf:   j.Workflow,
		}
		j.Root = ec
	}
	ec.ctx = ctx
	ec.span = span
	ec.runtime = jr
	defer ec.End()
	j.AttemptCount++

	return ec.run(j.Workflow, true)
}

// ExecContext contains mostly immutable state for the execution of a workflow
// or stage.
type ExecContext struct {
	Name        string
	Children    []*ExecContext
	MutatedVars []byte

	// Runtime information that is not persisted between runs.

	wf           Workflow
	ctx          context.Context
	span         *trace.Span
	runtime      job.Runtime
	rollbacks    []OpFN // rollback operations, in order encountered.
	fullRollback bool
}

// ID returns the job id.
func (ec *ExecContext) ID() string {
	return ec.runtime.ID()
}

// subContext creates a new child context.
func (ec *ExecContext) subContext(cfg *runOpts) *ExecContext {
	if cfg.name == "" {
		cfg.name = fmt.Sprintf("stage-%02d", len(ec.Children)+1)
	}
	ctx, span := trace.StartSpan(ec.ctx, "workflow.runner")
	span.AddAttributes(trace.StringAttribute("name", cfg.name))

	sc := &ExecContext{
		Name:         path.Join(ec.Name, cfg.name),
		wf:           ec.wf,
		ctx:          ctx,
		span:         span,
		runtime:      ec.runtime,
		fullRollback: ec.fullRollback,
	}
	ec.Children = append(ec.Children, sc)
	ec.rollbacks = append(ec.rollbacks, sc.revert)
	return sc
}

// End releases resources associated with the context.
func (ec *ExecContext) End() {
	ec.span.End()
	ec.span = nil
	ec.ctx = nil
}

// run a workflow.  If catchErrors is true, then internal.Error throws are
// caught and returned, otherwise errors are left uncaught and must be recovered
// by the caller.
func (ec *ExecContext) run(wf Workflow, catchErrors bool) error {
	ll := ec.Log()
	return backoff.Retry(func() error {
		ctx, span := trace.StartSpan(ec.ctx, "workflow.run")
		defer span.End()

		err := ec.once(ctx, wf, catchErrors)
		if err == nil {
			span.Annotate(nil, "success")
			return nil
		}

		span.Annotate([]trace.Attribute{
			trace.StringAttribute("err", err.Error())}, "failed")

		ll.WithError(err).Warn("workflow failed")
		if ec.fullRollback || job.IsUnrecoverable(err) {
			ec.revert()
		}

		if job.IsUnrecoverable(err) {
			ll.WithError(err).Error("unrecoverable error")
			return backoff.Permanent(err)
		}

		return err
	}, backoff.NewExponentialBackOff())
}

func (ec *ExecContext) once(ctx context.Context, wf Workflow, catchErrors bool) (err error) {
	panicked := true
	defer func() {
		if panicked && catchErrors {
			v := recover()
			if e, ok := v.(error); ok {
				err = e
			} else {
				err = errors.Errorf("internal error: %v", v)
			}
		}
	}()

	wf.Run(ctx, ec)
	panicked = false

	// Sync state on success.
	if upErr := ec.runtime.Update(ec.ctx, job.WithDescription(ec.Name)); upErr != nil {
		log.WithError(upErr).Error("unable to update job")
		err = upErr
	}

	return
}

func (ec *ExecContext) revert() {
	ec.Log().WithField("num", len(ec.rollbacks)).Debug("rolling back workflow")

	for i := len(ec.rollbacks) - 1; i >= 0; i-- {
		rollbackFN := ec.rollbacks[i]
		rollbackFN()
	}
}

// Log returns a logging channel for the current stage.
func (ec *ExecContext) Log() log.FieldLogger {
	return ec.runtime.Log().WithField("stage", ec.Name)
}

// Context returns the current context.
func (ec *ExecContext) Context() context.Context {
	return ec.ctx
}

// Do implements the Runtime interface.
func (ec *ExecContext) Do(fn OpFN, opts ...Opt) {
	cfg := &runOpts{}
	for _, o := range opts {
		o(cfg)
	}

	sc := ec.subContext(cfg)
	defer sc.End()
	sc.runStage(fn, cfg)
}

// OnRollback implements the Runtime interface.
func (ec *ExecContext) OnRollback(fn OpFN) {
	ec.Log().Debug("adding rollback function")
	ec.rollbacks = append(ec.rollbacks, fn)
}

// Embed implements the Runtime interface.
func (ec *ExecContext) Embed(wf Workflow, opts ...Opt) {
	cfg := &runOpts{}
	for _, o := range opts {
		o(cfg)
	}

	sc := ec.subContext(cfg)
	defer sc.End()
	err := sc.run(wf, false)
	if err != nil {
		ec.Log().WithError(err).Error("embedded workflow failed")
		ec.Fail(err)
	}
}

func (ec *ExecContext) runStage(fn OpFN, cfg *runOpts) {
	ll := ec.Log()
	ll.Debug("beginning stage")

	if ec.MutatedVars != nil {
		buf := bytes.NewBuffer(ec.MutatedVars)
		dec := gob.NewDecoder(buf)
		for i, out := range cfg.vars {
			if err := dec.Decode(out); err != nil {
				err = errors.Wrapf(err, "unable to deserialize var, index %d", i)
				ec.Fail(job.Unrecoverable(err))
			}
		}
		ll.Debug("stage cache reused")
		return
	}

	if fn != nil {
		fn()
		ll.Debug("stage function succeeded")
	}

	// Serialize output vars.
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	for _, out := range cfg.vars {
		if err := enc.Encode(out); err != nil {
			err = errors.Wrap(err, "unable to encode job")
			ec.Fail(job.Unrecoverable(err))
		}
	}
	ec.MutatedVars = buf.Bytes()

	// Always update state at successful completion of a step.
	// This ensures that we don't take multiple steps without recording progress.
	// Sync implements the Runtime interface.
	if err := ec.runtime.Update(ec.ctx, job.WithDescription(ec.Name)); err != nil {
		ec.Fail(err)
	}
	ll.Debug("stage results stored")
}

// Fail implements the Runtime interface.
func (ec *ExecContext) Fail(err error) {
	ec.span.Annotate([]trace.Attribute{
		trace.StringAttribute("error", err.Error()),
	}, "failed")
	panic(err)
}
