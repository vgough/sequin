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
)

func TestSingleStageWorkflow(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wf := &singleStageWorkflow{Count: 2}
	_, err := autotest.Run(ctx, t, "single_stage_test", wf)
	require.NoError(t, err)
	require.True(t, len(attempts) > 0)
}

var attempts = map[string]int{}

func init() {
	gob.Register(&singleStageWorkflow{})
}

type singleStageWorkflow struct {
	Count int
}

func (wf *singleStageWorkflow) Run(ctx context.Context, run workflow.Runtime) {
	for i := 0; i < wf.Count; i++ {
		n := attempts[run.ID()]
		attempts[run.ID()]++
		if n == 0 {
			run.Fail(errors.New("random fail"))
		}
	}
}
