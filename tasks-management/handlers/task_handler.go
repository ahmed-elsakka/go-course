package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-manager/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodPost:
		defer r.Body.Close()

		var body struct {
			Title string `json:"title"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		task, err := h.service.CreateTask(body.Title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(task)

	case http.MethodGet:
		tasks, err := h.service.GetTasks()
		if err != nil {
			http.Error(w, "failed to fetch tasks", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tasks)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid task ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodPut:
		defer r.Body.Close()

		var body struct {
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		task, err := h.service.UpdateTask(uint(id), body.Title, body.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(task)

	case http.MethodDelete:
		err := h.service.DeleteTask(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
