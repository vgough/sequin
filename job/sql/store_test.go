package sql

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vgough/sequin/job"
)

func TestCompatibility(t *testing.T) {
	job.TestCompatibility(t, TestStore(t))
}
