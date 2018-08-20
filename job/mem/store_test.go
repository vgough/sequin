package mem

import (
	"testing"

	"github.com/vgough/sequin/job"
)

func TestStore(t *testing.T) {
	job.TestCompatibility(t, NewStore())
}
