package main

import (
	"fmt"
	"os"

	"todo/internal/storage"
)

const defaultJSONFile = "todos.json"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	var store storage.Storage = storage.NewJSONStorage(defaultJSONFile)

	tl, err := store.Load()
	if err != nil {
		fmt.Println("error loading tasks:", err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		runAdd(&tl, os.Args[2:])
	case "list":
		runList(&tl, os.Args[2:])
		return
	case "done":
		runDone(&tl, os.Args[2:])
	case "delete":
		runDelete(&tl, os.Args[2:])
	case "-h", "--help", "help":
		printUsage()
		return
	default:
		fmt.Println("unknown command:", os.Args[1])
		printUsage()
		os.Exit(1)
	}

	if err := store.Save(tl); err != nil {
		fmt.Println("error saving tasks:", err)
		os.Exit(1)
	}

}

func printUsage() {
	fmt.Println("usage: todo <add|list|done|delete> [args]")
}
