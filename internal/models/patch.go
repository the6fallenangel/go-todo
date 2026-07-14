package models

import "time"

type TaskPatch struct {
	Description  *string
	Done         *bool
	Priority     *Priority
	DueDate      *time.Time
	ClearDueDate bool
	Tags         *[]string
}

func (t *Task) Apply(patch TaskPatch) {
	if patch.Description != nil {
		t.Description = *patch.Description
	}
	if patch.Done != nil {
		t.Done = *patch.Done
	}
	if patch.Priority != nil {
		t.Priority = *patch.Priority
	}
	if patch.ClearDueDate {
		t.DueDate = nil
	} else if patch.DueDate != nil {
		t.DueDate = patch.DueDate
	}
	if patch.Tags != nil {
		t.Tags = *patch.Tags
	}
}
