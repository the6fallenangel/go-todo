package storage

import (
	"path/filepath"
	"testing"

	"todo/internal/models"
)

func TestJSONStorage(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "todos.json")

	var store Storage = NewJSONStorage(path)

	var tl models.TaskList
	tl.Add("buy milk", models.PriorityLow, nil, nil)

	if err := store.Save(tl); err != nil {
		t.Fatalf("unexpected error saving: %v", err)
	}

	loaded, err := store.Load()
	if err != nil {
		t.Fatalf("unexpected error loading: %v", err)
	}
	if len(loaded.Tasks) != 1 {
		t.Fatalf("expected 1 task, got %d", len(loaded.Tasks))
	}
}
