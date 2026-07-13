package storage

import (
	"encoding/json"
	"os"
	"todo/internal/models"
)

const DefaultFile = "todos.json"

func Load(path string) (models.TaskList, error) {
	var tl models.TaskList
	data, err := os.ReadFile(path)
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

func Save(path string, tl models.TaskList) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
