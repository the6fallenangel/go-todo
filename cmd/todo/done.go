package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"todo/internal/storage"
)

func runDone(store storage.Storage, args []string) {
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
	if err := store.Done(id); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Println("marked task", id, "as done")
}
