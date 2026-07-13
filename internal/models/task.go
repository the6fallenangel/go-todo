package models

import (
	"time"
)

type Priority int

const (
	PriorityLow Priority = iota
	PriorityMedium
	PriorityHigh
)

func (p Priority) String() string {
	switch p {
	case PriorityHigh:
		return "high"
	case PriorityMedium:
		return "medium"
	default:
		return "low"
	}
}

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Done        bool       `json:"done"`
	Priority    Priority   `json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

func NewTask(id int, description string) Task {
	return Task{
		ID:          id,
		Description: description,
		Done:        false,
		Priority:    PriorityLow,
		CreatedAt:   time.Now(),
	}
}
