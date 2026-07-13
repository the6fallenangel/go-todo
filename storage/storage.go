package storage

import (
	"encoding/json"
	"os"
	"todo/models"
)

const defaultFile = "todos.json"

func Load() (models.TaskList, error) {
	var tl models.TaskList
	data, err := os.ReadFile(defaultFile)
	if os.IsNotExist(err) {
		return tl, nil
	}
	if err != nil {
		return tl, nil
	}
	if err := json.Unmarshal(data, &tl); err != nil {
		return tl, err
	}
	return tl, nil
}

func Save(tl models.TaskList) error {
	data, err := json.MarshalIndent(tl, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(defaultFile, data, 0644)
}
