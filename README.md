# go-todo

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go)
![CI](https://github.com/the6fallenangel/go-todo/actions/workflows/ci.yml/badge.svg)

A backend/CLI todo application written in Go, built as a project-based learning exercise.

Not a production app — the goal is to demonstrate solid, idiomatic Go across CLI, HTTP, and persistence layers while gradually evolving a real application.

## Features

- **CLI** (`cmd/todo`) and **REST API** (`cmd/todo-api`), sharing the same core logic
- **JSON and SQLite storage**, interchangeable through a `Storage` interface
- Backend selection through `.env` configuration
- Full CRUD via the API, including partial updates (`PATCH`) and query-param filtering
- Priority, due dates, tags
- Request logging and panic recovery middleware
- Graceful shutdown for the API server
- Unit and handler tests running against both storage implementations
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

`internal/` is enforced by the Go compiler itself, keeping these packages private to this module. Both `cmd` entry points depend only on the `Storage` interface — neither knows or cares whether data is stored in JSON or SQLite.

## Usage

### CLI

```bash
./todo add -priority high -due 2026-08-01 -tag work "ship feature"
./todo list -tag work
./todo done 1
./todo delete 2
```

### REST API

```bash
curl -X POST localhost:8080/tasks \
  -d '{"description":"ship feature","priority":"high","tags":["work"]}'

curl localhost:8080/tasks
curl localhost:8080/tasks/1
curl 'localhost:8080/tasks?tag=work'
curl 'localhost:8080/tasks?done=true'

curl -X PATCH localhost:8080/tasks/1 -d '{"done":true}'
curl -X DELETE localhost:8080/tasks/1
```

### Configuration

```bash
cp .env.example .env
```

```env
BACKEND=json # or sqlite
JSON_PATH=todos.json
SQLITE_PATH=todos.db
```

## Testing

```bash
go test ./... -v
```

The same storage test suite runs against both JSON and SQLite implementations, ensuring both satisfy the same `Storage` contract.

## Status

This project served as a hands-on introduction to Go — structs and interfaces, error handling, `net/http`, SQL, and testing. It's complete for its original purpose and not under active feature development.
