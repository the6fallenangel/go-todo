package main

import (
	"fmt"
	"log"
	"net/http"
	"todo/internal/config"
	"todo/internal/storage"
)

func main() {
	cfg := config.Load()

	store, err := newStorage(cfg)
	if err != nil {
		log.Fatal("error initializing storage:", err)
	}

	mux := http.NewServeMux()
	registerRoutes(mux, store)

	addr := ":8080"
	fmt.Println("listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func newStorage(cfg config.Config) (storage.Storage, error) {
	switch cfg.Backend {
	case "json":
		return storage.NewJSONStorage(cfg.JSONPath), nil
	case "sqlite":
		return storage.NewSQLiteStorage(cfg.SQLitePath)
	default:
		return nil, fmt.Errorf("unknown backend %q in .env (want json or sqlite)", cfg.Backend)
	}
}
