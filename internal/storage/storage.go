package storage

import (
	"todo/internal/models"
)

type Storage interface {
	List() ([]models.Task, error)
	Add(task models.Task) (models.Task, error)
	Done(id int) error
	Delete(id int) error
}
