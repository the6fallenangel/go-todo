package models

import "testing"

func TestAdd(t *testing.T) {
	var tl TaskList

	task := tl.Add("buy milk", PriorityLow, nil, nil)

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
	if task.Description != "buy milk" {
		t.Errorf("expected description %q, got %q", "buy milk", task.Description)
	}
	if len(tl.Tasks) != 1 {
		t.Errorf("expected 1 task in list, got %d", len(tl.Tasks))
	}
}

func TestAddIncrementsNextID(t *testing.T) {
	var tl TaskList

	tl.Add("a", PriorityLow, nil, nil)
	tl.Add("b", PriorityLow, nil, nil)
	task := tl.Add("c", PriorityLow, nil, nil)

	if task.ID != 3 {
		t.Errorf("expected ID 3, got %d", task.ID)
	}
}

func TestAddIDsSurviveDelete(t *testing.T) {
	var tl TaskList

	tl.Add("a", PriorityLow, nil, nil) // ID 1
	tl.Add("b", PriorityLow, nil, nil) // ID 2

	if err := tl.Delete(1); err != nil {
		t.Fatalf("unexpected error deleting: %v", err)
	}

	task := tl.Add("c", PriorityLow, nil, nil)
	if task.ID == 2 {
		t.Errorf("new task collided with existing ID 2")
	}
}

func TestDone(t *testing.T) {
	var tl TaskList
	tl.Add("buy milk", PriorityLow, nil, nil)

	if err := tl.Done(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !tl.Tasks[0].Done {
		t.Errorf("expected task to be marked done")
	}
}

func TestDoneNotFound(t *testing.T) {
	var tl TaskList

	err := tl.Done(99)
	if err == nil {
		t.Errorf("expected error for missing task, got nil")
	}
}

func TestDelete(t *testing.T) {
	var tl TaskList
	tl.Add("a", PriorityLow, nil, nil)
	tl.Add("b", PriorityLow, nil, nil)

	if err := tl.Delete(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tl.Tasks) != 1 {
		t.Errorf("expected 1 task remaining, got %d", len(tl.Tasks))
	}
	if tl.Tasks[0].ID != 2 {
		t.Errorf("expected remaining task to have ID 2, got %d", tl.Tasks[0].ID)
	}
}

func TestDeleteNotFound(t *testing.T) {
	var tl TaskList

	err := tl.Delete(99)
	if err == nil {
		t.Errorf("expected error for missing task, got nil")
	}
}
