package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"github.com/the6fallenangel/go-todo/internal/storage"
)

func runList(store storage.Storage, args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	doneOnly := fs.Bool("done", false, "show only completed tasks")
	pendingOnly := fs.Bool("pending", false, "show only pending tasks")
	tagFilter := fs.String("tag", "", "show only tasks with this tag")
	fs.Parse(args)

	tasks, err := store.List()
	if err != nil {
		fmt.Println("error listing tasks:", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks.")
		return
	}

	for _, t := range tasks {
		if *doneOnly && !t.Done {
			continue
		}
		if *pendingOnly && t.Done {
			continue
		}
		if *tagFilter != "" && !slices.Contains(t.Tags, *tagFilter) {
			continue
		}
		status := " "
		if t.Done {
			status = "x"
		}
		due := ""
		if t.DueDate != nil {
			due = " (due " + t.DueDate.Format("2006-01-02") + ")"
		}
		tagsStr := ""
		if len(t.Tags) > 0 {
			tagsStr = fmt.Sprintf(" %v", t.Tags)
		}
		fmt.Printf("[%s] #%d: %s [%s]%s%s\n", status, t.ID, t.Description, t.Priority, due, tagsStr)
	}
}
