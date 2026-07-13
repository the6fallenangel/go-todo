package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/internal/models"
	"todo/internal/storage"
)

func handleList(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := store.List()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}
		writeJSON(w, http.StatusOK, tasks)
	}
}

type AddTaskRequest struct {
	Description string   `json:"description"`
	Priority    string   `json:"priority"`
	Tags        []string `json:"tags"`
}

func handleAdd(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddTaskRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}
		if req.Description == "" {
			writeError(w, http.StatusBadRequest, errString("description is required"))
			return
		}

		priority := models.PriorityLow
		switch req.Priority {
		case "medium":
			priority = models.PriorityMedium
		case "high":
			priority = models.PriorityHigh
		}

		task := models.NewTask(req.Description)
		task.Priority = priority
		task.Tags = req.Tags

		saved, err := store.Add(task)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}
		writeJSON(w, http.StatusCreated, saved)
	}
}

func handleDone(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, errString("invalid task id"))
			return
		}
		if err := store.Done(id); err != nil {
			writeError(w, http.StatusNotFound, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func handleDelete(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, errString("invalid task id"))
			return
		}
		if err := store.Delete(id); err != nil {
			writeError(w, http.StatusNotFound, err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
