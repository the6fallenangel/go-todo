# go-todo

A CLI todo application written in Go, built as a learning project — each feature is added deliberately to explore a specific Go concept, with a clean commit history to track progress.

## Features

- Add, list, complete, and delete tasks
- Priority levels (low / medium / high)
- Due dates
- Tags with filtering (`-tag`, `-done`, `-pending`)
- JSON persistence (`todos.json`)
- Subcommand-based CLI using the standard `flag` package
- Standard Go project layout (`cmd/`, `internal/`)

## Usage

```bash
go build -o todo ./cmd/todo

./todo add -priority high -due 2026-08-01 -tag work -tag urgent "ship feature"
./todo list
./todo list -tag work
./todo list -pending
./todo done 1
./todo delete 2
```
