package storage

import (
	"path/filepath"
	"testing"

	"github.com/the6fallenangel/go-todo/internal/models"
)

func TestStorageImplementations(t *testing.T) {
	backends := []struct {
		name string
		new  func(t *testing.T) Storage
	}{
		{
			name: "json",
			new: func(t *testing.T) Storage {
				dir := t.TempDir()
				return NewJSONStorage(filepath.Join(dir, "todos.json"))
			},
		},
		{
			name: "sqlite",
			new: func(t *testing.T) Storage {
				dir := t.TempDir()
				s, err := NewSQLiteStorage(filepath.Join(dir, "todos.db"))
				if err != nil {
					t.Fatalf("creating sqlite storage: %v", err)
				}
				return s
			},
		},
	}

	for _, b := range backends {
		t.Run(b.name, func(t *testing.T) {
			store := b.new(t)

			saved, err := store.Add(models.NewTask("buy milk"))
			if err != nil {
				t.Fatalf("unexpected error adding: %v", err)
			}
			if saved.ID == 0 {
				t.Errorf("expected non-zero id")
			}

			tasks, err := store.List()
			if err != nil {
				t.Fatalf("unexpected error listing: %v", err)
			}
			if len(tasks) != 1 {
				t.Fatalf("expected 1 task, got %d", len(tasks))
			}

			if err := store.Done(saved.ID); err != nil {
				t.Fatalf("unexpected error marking done: %v", err)
			}
			tasks, _ = store.List()
			if !tasks[0].Done {
				t.Errorf("expected task to be marked done")
			}

			if err := store.Delete(saved.ID); err != nil {
				t.Fatalf("unexpected error deleting: %v", err)
			}
			tasks, _ = store.List()
			if len(tasks) != 0 {
				t.Errorf("expected 0 tasks after delete, got %d", len(tasks))
			}

			if err := store.Done(999); err == nil {
				t.Errorf("expected error marking nonexistent task done")
			}
			if err := store.Delete(999); err == nil {
				t.Errorf("expected error deleting nonexistent task")
			}
		})
	}
}
