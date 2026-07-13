package storage

import (
	"path/filepath"
	"testing"

	"todo/internal/models"
)

func TestSaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "todos.json")

	var tl models.TaskList
	tl.Add("buy milk", models.PriorityLow, nil, nil)
	tl.Add("walk dog", models.PriorityHigh, nil, nil)

	if err := Save(path, tl); err != nil {
		t.Fatalf("unexpected error saving: %v", err)
	}

	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("unexpected error loading: %v", err)
	}

	if len(loaded.Tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(loaded.Tasks))
	}
	if loaded.Tasks[0].Description != "buy milk" {
		t.Errorf("expected first task %q, got %q", "buy milk", loaded.Tasks[0].Description)
	}
	if loaded.NextID != tl.NextID {
		t.Errorf("expected NextID %d, got %d", tl.NextID, loaded.NextID)
	}
}

func TestLoadMissingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "does-not-exist.json")

	tl, err := Load(path)
	if err != nil {
		t.Fatalf("expected no error for missing file, got %v", err)
	}
	if len(tl.Tasks) != 0 {
		t.Errorf("expected empty task list, got %d tasks", len(tl.Tasks))
	}
}
