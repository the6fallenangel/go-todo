package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/the6fallenangel/go-todo/internal/models"
	"github.com/the6fallenangel/go-todo/internal/storage"
)

func handleList(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tasks, err := store.List()
		if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		q := r.URL.Query()
		if doneParam := q.Get("done"); doneParam != "" {
			want := doneParam == "true"
			tasks = filterTasks(tasks, func(t models.Task) bool { return t.Done == want })
		}
		if tag := q.Get("tag"); tag != "" {
			tasks = filterTasks(tasks, func(t models.Task) bool { return slices.Contains(t.Tags, tag) })
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

func handleGet(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, errString("invalid task id"))
			return
		}
		task, err := store.Get(id)
		if err != nil {
			writeError(w, http.StatusNotFound, err)
			return
		}
		writeJSON(w, http.StatusOK, task)
	}
}

type updateTaskRequest struct {
	Description *string   `json:"description"`
	Done        *bool     `json:"done"`
	Priority    *string   `json:"priority"`
	DueDate     *string   `json:"due_date"`
	Tags        *[]string `json:"tags"`
}

func handleUpdate(store storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, errString("invalid task id"))
			return
		}

		var req updateTaskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}

		patch := models.TaskPatch{
			Description: req.Description,
			Done:        req.Done,
			Tags:        req.Tags,
		}

		if req.Priority != nil {
			p := models.PriorityLow
			switch *req.Priority {
			case "medium":
				p = models.PriorityMedium
			case "high":
				p = models.PriorityHigh
			}
			patch.Priority = &p
		}

		if req.DueDate != nil {
			if *req.DueDate == "" {
				patch.ClearDueDate = true
			} else {
				parsed, err := time.Parse("2006-01-02", *req.DueDate)
				if err != nil {
					writeError(w, http.StatusBadRequest, errString("invalid due_date, expected YYYY-MM-DD"))
					return
				}
				patch.DueDate = &parsed
			}
		}

		updated, err := store.Update(id, patch)
		if err != nil {
			writeError(w, http.StatusNotFound, err)
			return
		}
		writeJSON(w, http.StatusOK, updated)
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

func filterTasks(tasks []models.Task, keep func(models.Task) bool) []models.Task {
	filtered := make([]models.Task, 0, len(tasks))
	for _, t := range tasks {
		if keep(t) {
			filtered = append(filtered, t)
		}
	}
	return filtered
}
