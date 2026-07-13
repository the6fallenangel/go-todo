package main

import (
	"fmt"
	"os"

	"todo/internal/config"
	"todo/internal/storage"
)

const defaultJSONFile = "todos.json"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cfg := config.Load()
	store, err := newStorage(cfg)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		runAdd(store, os.Args[2:])
	case "list":
		runList(store, os.Args[2:])
		return
	case "done":
		runDone(store, os.Args[2:])
	case "delete":
		runDelete(store, os.Args[2:])
	case "-h", "--help", "help":
		printUsage()
		return
	default:
		fmt.Println("unknown command:", os.Args[1])
		printUsage()
		os.Exit(1)
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

func printUsage() {
	fmt.Println("usage: todo <add|list|done|delete> [args]")
	fmt.Println("configure storage backend via .env (see .env.example)")
}
