package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"task-management/services"
)

type TaskHandler struct {
	service *services.TaskService
}

func NewTaskHandler(service *services.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		defer r.Body.Close()

		var requestJson struct {
			Title string `json:"title"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestJson)

		if err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		task, err := h.service.CreateTask(requestJson.Title)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(task)

	case http.MethodGet:
		tasks, err := h.service.GetAll()

		if err != nil {
			http.Error(w, "Error fetching the tasks", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(tasks)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) HandleTaskById(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid task Id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		defer r.Body.Close()

		var requestJson struct {
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestJson)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		task, err := h.service.UpdateTask(uint(id), requestJson.Title, requestJson.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(task)

	case http.MethodDelete:
		err = h.service.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
