package storage

import (
	"encoding/json"
	"os"
	"todo/internal/models"
)

type JSONStorage struct {
	Path string
}

func NewJSONStorage(path string) *JSONStorage {
	return &JSONStorage{Path: path}
}

func (s *JSONStorage) Load() (models.TaskList, error) {
	var tl models.TaskList
	data, err := os.ReadFile(s.Path)
	if os.IsNotExist(err) {
		return tl, nil
	}
	if err != nil {
		return tl, err
	}
	if err := json.Unmarshal(data, &tl); err != nil {
		return tl, err
	}
	return tl, nil
}

func (s *JSONStorage) Save(tl models.TaskList) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Path, data, 0644)
}
