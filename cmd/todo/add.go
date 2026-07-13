package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"todo/internal/models"
)

func runAdd(tl *models.TaskList, args []string) {
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

	task := tl.Add(fs.Arg(0), priority, due, tags)
	fmt.Printf("added task #%d: %s\n", task.ID, task.Description)
}
