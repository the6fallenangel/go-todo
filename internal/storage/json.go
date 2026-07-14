package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/the6fallenangel/go-todo/internal/models"
)

type jsonFile struct {
	Tasks  []models.Task `json:"tasks"`
	NextID int           `json:"next_id"`
}

type JSONStorage struct {
	Path string
}

func NewJSONStorage(path string) *JSONStorage {
	return &JSONStorage{Path: path}
}

func (s *JSONStorage) read() (jsonFile, error) {
	var f jsonFile
	data, err := os.ReadFile(s.Path)
	if os.IsNotExist(err) {
		return f, nil
	}
	if err != nil {
		return f, err
	}
	if err := json.Unmarshal(data, &f); err != nil {
		return f, err
	}
	return f, nil
}

func (s *JSONStorage) write(f jsonFile) error {
	data, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, data, 0644)
}

func (s *JSONStorage) List() ([]models.Task, error) {
	f, err := s.read()
	if err != nil {
		return nil, err
	}
	return f.Tasks, nil
}

func (s *JSONStorage) Add(task models.Task) (models.Task, error) {
	f, err := s.read()
	if err != nil {
		return models.Task{}, err
	}
	f.NextID++
	task.ID = f.NextID
	f.Tasks = append(f.Tasks, task)
	if err := s.write(f); err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (s *JSONStorage) Get(id int) (models.Task, error) {
	f, err := s.read()
	if err != nil {
		return models.Task{}, err
	}
	for _, t := range f.Tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return models.Task{}, fmt.Errorf("task with id %d not found", id)
}

func (s *JSONStorage) Update(id int, patch models.TaskPatch) (models.Task, error) {
	f, err := s.read()
	if err != nil {
		return models.Task{}, err
	}
	for i := range f.Tasks {
		if f.Tasks[i].ID == id {
			f.Tasks[i].Apply(patch)
			if err := s.write(f); err != nil {
				return models.Task{}, err
			}
			return f.Tasks[i], nil
		}
	}
	return models.Task{}, fmt.Errorf("task with id %d not found", id)
}
func (s *JSONStorage) Done(id int) error {
	f, err := s.read()
	if err != nil {
		return err
	}
	for i := range f.Tasks {
		if f.Tasks[i].ID == id {
			f.Tasks[i].Done = true
			return s.write(f)
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}

func (s *JSONStorage) Delete(id int) error {
	f, err := s.read()
	if err != nil {
		return err
	}
	for i := range f.Tasks {
		if f.Tasks[i].ID == id {
			f.Tasks = append(f.Tasks[:i], f.Tasks[i+1:]...)
			return s.write(f)
		}
	}
	return fmt.Errorf("task with id %d not found", id)
}
