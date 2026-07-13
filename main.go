package main

import (
	"fmt"
	"os"
	"strconv"
	"todo/storage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: todo <add|list|done|delete> [args]")
		os.Exit(1)
	}
	tl, err := storage.Load()
	if err != nil {
		fmt.Println("error loading tasks:", err)
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("usage: todo add \"description\"")
			os.Exit(1)
		}
		task := tl.Add(os.Args[2])
		fmt.Printf("added task #%d: %s\n", task.ID, task.Description)
	case "list":
		if len(tl.Tasks) == 0 {
			fmt.Println("No tasks.")
			return
		}
		for _, t := range tl.Tasks {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] #%d: %s\n", status, t.ID, t.Description)
		}
		return
	case "done":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id:", os.Args[2])
			os.Exit(1)
		}
		if err := tl.Complete(id); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println("marked task", id, "as done")
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id:", os.Args[2])
			os.Exit(1)
		}
		if err := tl.Delete(id); err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Println("deleted task", id)
	default:
		fmt.Println("unknown command:", command)
		os.Exit(1)
	}
	if err := storage.Save(tl); err != nil {
		fmt.Println("error saving tasks:", err)
		os.Exit(1)
	}

}
