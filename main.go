package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"todo/models"
	"todo/storage"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	tl, err := storage.Load()
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

	if err := storage.Save(tl); err != nil {
		fmt.Println("error saving tasks:", err)
		os.Exit(1)
	}

}

func printUsage() {
	fmt.Println("usage: todo <add|list|done|delete> [args]")
}

func runAdd(tl *models.TaskList, args []string) {
	fs := flag.NewFlagSet("add", flag.ExitOnError)
	priorityFlag := fs.String("priority", "low", "priority: low, medium, high")
	dueFlag := fs.String("due", "", "due date, format YYYY-MM-DD")
	fs.Parse(args)

	if fs.NArg() < 1 {
		fmt.Println(`usage: todo add [-priority low|medium|high] [-due YYYY-MM-DD] "description"`)
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

	task := tl.Add(fs.Arg(0), priority, due)
	fmt.Printf("added task #%d: %s\n", task.ID, task.Description)
}

func runList(tl *models.TaskList, args []string) {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	doneOnly := fs.Bool("done", false, "show only completed tasks")
	pendingOnly := fs.Bool("pending", false, "show only pending tasks")
	fs.Parse(args)

	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks.")
		return
	}

	for _, t := range tl.Tasks {
		if *doneOnly && !t.Done {
			continue
		}
		if *pendingOnly && t.Done {
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
		fmt.Printf("[%s] #%d: %s [%s]%s\n", status, t.ID, t.Description, t.Priority, due)
	}
}

func runDone(tl *models.TaskList, args []string) {
	fs := flag.NewFlagSet("done", flag.ExitOnError)
	fs.Parse(args)

	if fs.NArg() < 1 {
		fmt.Println("usage: todo done <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(fs.Arg(0))
	if err != nil {
		fmt.Println("invalid id:", fs.Arg(0))
		os.Exit(1)
	}

	if err := tl.Done(id); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("marked task", id, "as done")
}

func runDelete(tl *models.TaskList, args []string) {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	fs.Parse(args)
	if fs.NArg() < 1 {
		fmt.Println("usage: todo delete <id>")
		os.Exit(1)
	}

	id, err := strconv.Atoi(fs.Arg(0))
	if err != nil {
		fmt.Println("invalid id:", fs.Arg(0))
		os.Exit(1)
	}

	if err := tl.Delete(id); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("deleted task", id)
}
