package storage

import (
	"fmt"

	"github.com/the6fallenangel/go-todo/internal/config"
)

func New(cfg config.Config) (Storage, error) {
	switch cfg.Backend {
	case "json":
		return NewJSONStorage(cfg.JSONPath), nil
	case "sqlite":
		return NewSQLiteStorage(cfg.SQLitePath)
	default:
		return nil, fmt.Errorf("unknown backend %q in .env (want json or sqlite)", cfg.Backend)
	}
}
