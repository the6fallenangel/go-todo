package storage

import (
	"github.com/the6fallenangel/go-todo/internal/models"
)

type Storage interface {
	List() ([]models.Task, error)
	Add(task models.Task) (models.Task, error)
	Get(id int) (models.Task, error)
	Update(id int, patch models.TaskPatch) (models.Task, error)
	Done(id int) error
	Delete(id int) error
}
