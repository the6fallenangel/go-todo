package models

import "fmt"

type TaskList struct {
	Tasks  []Task `json:"tasks"`
	NextID int    `json:"next_id"`
}

func (tl *TaskList) Add(description string) Task {
	tl.NextID++
	task := NewTask(tl.NextID, description)
	tl.Tasks = append(tl.Tasks, task)
	return task
}

func (tl *TaskList) Complete(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found.", id)
}

func (tl *TaskList) Delete(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with id %d not found.", id)
}
