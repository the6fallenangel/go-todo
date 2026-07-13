package storage

import (
	"todo/internal/models"
)

type Storage interface {
	Load() (models.TaskList, error)
	Save(models.TaskList) error
}
