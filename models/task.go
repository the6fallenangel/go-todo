package models

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewTask(id int, description string) Task {
	return Task{
		ID:          id,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
	}
}
