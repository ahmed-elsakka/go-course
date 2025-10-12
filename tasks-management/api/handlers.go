package api

import (
	"encoding/json"
	"net/http"
	"tasks-management/models"
	"time"
)

type Handler struct {
	Repo models.TaskRepository
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	json.NewDecoder(r.Body).Decode(&t)
	t.Status = "Pending"
	t.CreatedAt = time.Now()
	h.Repo.Create(&t)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}

func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, _ := h.Repo.GetAll(1)
	json.NewEncoder(w).Encode(tasks)
}
