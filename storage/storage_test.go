package storage

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/vgough/sequin/storage/ent"
	"github.com/vgough/sequin/storage/ent/enttest"
	"github.com/vgough/sequin/storage/ent/migrate"
)

func TestXXX(t *testing.T) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)
	defer client.Close()
	// ...
}
