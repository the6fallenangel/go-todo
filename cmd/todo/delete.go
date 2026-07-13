package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"todo/internal/models"
)

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
