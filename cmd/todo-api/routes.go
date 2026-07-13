package main

import (
	"net/http"
	"todo/internal/storage"
)

func registerRoutes(mux *http.ServeMux, store storage.Storage) {
	mux.HandleFunc("GET /tasks", handleList(store))
	mux.HandleFunc("POST /tasks", handleAdd(store))
	mux.HandleFunc("POST /tasks/{id}/done", handleDone(store))
	mux.HandleFunc("DELETE /tasks/{id}", handleDelete(store))
}
