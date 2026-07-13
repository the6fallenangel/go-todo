package storage

import (
	"fmt"
	"strings"
	"time"
	"todo/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteTaskRow struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Description string
	Done        bool
	Priority    int
	DueDate     *time.Time
	Tags        string
	CreatedAt   time.Time
}

func (sqliteTaskRow) TableName() string { return "tasks" }

type SQLiteStorage struct {
	db *gorm.DB
}

func NewSQLiteStorage(path string) (*SQLiteStorage, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("opening sqlite db: %w", err)
	}
	if err := db.AutoMigrate(&sqliteTaskRow{}); err != nil {
		return nil, fmt.Errorf("migrating schema: %w", err)
	}
	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) List() ([]models.Task, error) {
	var rows []sqliteTaskRow

	if err := s.db.Order("id").Find(&rows).Error; err != nil {
		return nil, fmt.Errorf("listing tasks: %w", err)
	}

	tasks := make([]models.Task, 0, len(rows))
	for _, r := range rows {
		tasks = append(tasks, rowToTask(r))
	}

	return tasks, nil
}

func (s *SQLiteStorage) Add(task models.Task) (models.Task, error) {
	row := taskToRow(task)
	if err := s.db.Create(&row).Error; err != nil {
		return models.Task{}, fmt.Errorf("inserting task: %w", err)
	}
	return rowToTask(row), nil
}

func (s *SQLiteStorage) Done(id int) error {
	result := s.db.Model(&sqliteTaskRow{}).Where("id = ?", id).Update("done", true)
	if result.Error != nil {
		return fmt.Errorf("marking task done: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}
	return nil
}

func (s *SQLiteStorage) Delete(id int) error {
	result := s.db.Delete(&sqliteTaskRow{}, id)
	if result.Error != nil {
		return fmt.Errorf("deleting task: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("task with id %d not found", id)
	}
	return nil
}

func rowToTask(r sqliteTaskRow) models.Task {
	t := models.Task{
		ID:          int(r.ID),
		Description: r.Description,
		Done:        r.Done,
		Priority:    models.Priority(r.Priority),
		DueDate:     r.DueDate,
		CreatedAt:   r.CreatedAt,
	}
	if r.Tags != "" {
		t.Tags = strings.Split(r.Tags, ",")
	}
	return t
}

func taskToRow(t models.Task) sqliteTaskRow {
	return sqliteTaskRow{
		Description: t.Description,
		Done:        t.Done,
		Priority:    int(t.Priority),
		DueDate:     t.DueDate,
		Tags:        strings.Join(t.Tags, ","),
		CreatedAt:   t.CreatedAt,
	}
}
