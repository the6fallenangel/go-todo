package models

import "fmt"

type TaskList struct {
	Tasks []Task
}

func (tl *TaskList) Add(description string) Task {
	id := len(tl.Tasks)
	task := NewTask(id, description)
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
	return fmt.Errorf("Task with id %d not found.", id)
}

func (tl *TaskList) Delete(id int) error {
	for i := range tl.Tasks {
		if tl.Tasks[i].ID == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Task with id %d not found.", id)
}
