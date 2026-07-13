package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"github.com/the6fallenangel/go-todo/internal/models"
	"github.com/the6fallenangel/go-todo/internal/storage"
)

func runAdd(store storage.Storage, args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	priorityFlag := fs.String("priority", "low", "priority: low, medium, high")
	dueFlag := fs.String("due", "", "due date, format YYYY-MM-DD")

	var tags tagList
	fs.Var(&tags, "tag", "tag (repeatable), e.g. -tag work -tag urgent")

	fs.Parse(args)

	if fs.NArg() < 1 {
		fmt.Println(`usage: todo add [-priority low|medium|high] [-due YYYY-MM-DD] [-tag name]... "description"`)
		os.Exit(1)
	}

	priority := models.PriorityLow

	switch *priorityFlag {
	case "medium":
		priority = models.PriorityMedium
	case "high":
		priority = models.PriorityHigh
	}

	var due *time.Time

	if *dueFlag != "" {
		parsed, err := time.Parse("2006-01-02", *dueFlag)
		if err != nil {
			fmt.Println("invalid due date, expected YYYY-MM-DD:", *dueFlag)
			os.Exit(1)
		}
		due = &parsed
	}

	task := models.NewTask(fs.Arg(0))
	task.Priority = priority
	task.DueDate = due
	task.Tags = tags

	saved, err := store.Add(task)
	if err != nil {
		fmt.Println("error adding task:", err)
		os.Exit(1)
	}
	fmt.Printf("added task #%d: %s\n", saved.ID, saved.Description)
}
