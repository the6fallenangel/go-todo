# go-todo

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=go)
![CI](https://github.com/the6fallenangel/go-todo/actions/workflows/ci.yml/badge.svg)

A backend/CLI todo application written in Go, built as a project-based learning exercise.

Not a production app — the goal is to demonstrate solid, idiomatic Go across CLI, HTTP, and persistence layers while gradually evolving a real application.

## Features

- **CLI** (`cmd/todo`) and **REST API** (`cmd/todo-api`) sharing the same core logic
- **JSON and SQLite storage**, interchangeable through a `Storage` interface
- Backend selection through `.env` configuration
- Priority, due dates, tags, and filtering
- Graceful shutdown support for the API server
- Unit tests running against both storage implementations
- CI via GitHub Actions (build, vet, test)

## Project layout

```
.
├── cmd
│   ├── todo
│   │   └── CLI application
│   │
│   └── todo-api
│       └── REST API server (net/http)
│
├── internal
│   ├── models
│   │   └── Task domain types
│   │
│   ├── storage
│   │   └── Storage interface + JSON/SQLite implementations
│   │
│   └── config
│       └── .env configuration loading
│
├── go.mod
├── go.sum
└── README.md
```

`internal/` is enforced by the Go compiler itself, keeping these packages private to this module.

Both `cmd` entry points depend only on the `Storage` interface. They do not know or care whether data is stored in JSON or SQLite.

## Architecture

```
              CLI
               |
               |
       Storage Interface
               |
        ----------------
        |              |
       JSON          SQLite
```

The storage backend can be changed through configuration without modifying application logic.

## Usage

### CLI

```bash
./todo add -priority high -due 2026-08-01 -tag work "ship feature"

./todo list -tag work

./todo done 1

./todo delete 2
```

### REST API

Example requests:

```bash
curl -X POST localhost:8080/tasks \
-d '{"description":"ship feature","priority":"high"}'

curl localhost:8080/tasks

curl -X POST localhost:8080/tasks/1/done

curl -X DELETE localhost:8080/tasks/1
```

### Configuration

```bash
cp .env.example .env
```

Example:

```env
BACKEND=json (or sqlite)
JSON_PATH=todos.json
SQLITE_PATH=todos.db
```

## Testing

```bash
go test ./... -v
```

The same storage test suite runs against both JSON and SQLite implementations, ensuring both satisfy the same `Storage` contract.
