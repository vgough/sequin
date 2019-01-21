package autotest

import (
	"bytes"
	"context"
	"encoding/gob"
	"net/http"
	_ "net/http/pprof" // profiling
	"os"
	"sync"

	"contrib.go.opencensus.io/exporter/stackdriver"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // For test db.
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"github.com/vgough/sequin/job"
	"github.com/vgough/sequin/job/sql"
	"github.com/vgough/sequin/workflow"
	"go.opencensus.io/trace"
)

// Testing is the minimum amount of testing.T needed for running tests.
// This avoid pulling in the testing package during normal builds.
type Testing interface {
	require.TestingT
	Logf(fmt string, args ...interface{})
	Fatalf(fmt string, args ...interface{})
}

var testInit sync.Once

// Run automated tests on a workflow to check invariants.
func Run(ctx context.Context, t Testing, name string, wf workflow.Workflow) (workflow.Workflow, error) {
	testInit.Do(func() {
		log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		go func() {
			err := http.ListenAndServe("localhost:6060", nil)
			log.Info().Err(err).Msg("http server exited")
		}()
	})

	// Enable the Stackdriver Tracing exporter
	gid := os.Getenv("GCP_PROJECTID")
	if gid != "" {
		log.Info().Msg("enabling trace exporter")
		sd, err := stackdriver.NewExporter(stackdriver.Options{
			ProjectID: gid,
		})
		if err != nil {
			t.Fatalf("Failed to create the Stackdriver exporter: %v", err)
		}
		defer sd.Flush()

		// Register/enable the trace exporter
		trace.RegisterExporter(sd)

		// For demo purposes, set the trace sampling probability to be high
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(1.0)})
	}

	ctx, span := trace.StartSpan(ctx, "autotest")
	defer span.End()
	span.AddAttributes(trace.StringAttribute("test_name", name))

	// Serialize workflow state.
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&wf)
	require.NoError(t, err, "unable to encode workflow")
	twf := &testWorkflow{t: t, Workflow: wf, Initial: buf.Bytes()}

	sc := sql.TestStore(t)
	q := sql.TestQueue(t)

	sctx, cancel := context.WithCancel(ctx)
	sch := job.NewScheduler(q, sc)

	// Start background worker.
	go sch.WorkerLoop(sctx)

	// Start work.
	wfjob := workflow.NewRunnable(name, twf)
	rec, err := sch.Enqueue(sctx, wfjob, "test="+name)
	require.NoError(t, err, "unable to schedule workflow")

	snap, err := sch.WaitForCompletion(sctx, rec)
	require.NoError(t, err, "workflow failed")
	require.True(t, snap.IsFinal(), "workflow not finalized")

	// Load from last state.
	run, err := snap.Runnable()
	require.NoError(t, err)
	wf = run.(*workflow.JobWrapper).Workflow.(*testWorkflow).Workflow

	cancel()
	sch.Close()

	if snap.Result() != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeUnknown,
			Message: snap.Result().Error(),
		})
	}

	return wf, snap.Result()
}

// testWorkflow implements a workflow testing runtime.
// The testing is implemented as (guess what) - a workflow!
type testWorkflow struct {
	t        Testing
	Workflow workflow.Workflow
	Initial  []byte // Initial workflow state in serialized form.  For restarting tests.
}

var _ workflow.Workflow = &testWorkflow{}

func init() {
	gob.Register(&testWorkflow{})
}

// newInstance creates a new Workflow instance from the serialized state.
func (twf *testWorkflow) newInstance(rt workflow.Runtime) workflow.Workflow {
	dec := gob.NewDecoder(bytes.NewBuffer(twf.Initial))
	var decoded workflow.Workflow
	if err := dec.Decode(&decoded); err != nil {
		rt.Fail(job.Unrecoverable(errors.Wrap(err, "unable to decode workflow")))
	}
	return decoded
}

// Run contains the test workflow.
func (twf *testWorkflow) Run(ctx context.Context, rt workflow.Runtime) {
	// Start the workflow to be tested, using the instrumented test runtime.
	// testRT := &testRuntime{
	// 	t:         twf.t,
	// 	base:      rt,
	// 	nextSteps: list.New(),
	// }

	rt.Embed(&testConsistency{twf: twf}, workflow.Name("consistency"))
}

func (twf *testWorkflow) Result() error {
	return nil
}
