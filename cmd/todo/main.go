package main

import (
	"fmt"
	"os"

	"github.com/the6fallenangel/go-todo/internal/config"
	"github.com/the6fallenangel/go-todo/internal/storage"
)

const defaultJSONFile = "todos.json"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cfg := config.Load()
	store, err := storage.New(cfg)
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

func printUsage() {
	fmt.Println("usage: todo <add|list|done|delete> [args]")
	fmt.Println("configure storage backend via .env (see .env.example)")
}
