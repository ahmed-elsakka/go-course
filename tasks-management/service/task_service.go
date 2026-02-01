package service

import (
	"errors"
	"strings"
	"task-manager/models"
	"task-manager/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string) (*models.Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	task := &models.Task{
		Title:     title,
		Completed: false,
	}

	err := s.repo.Create(task)
	return task, err
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) UpdateTask(id uint, title string, completed bool) (*models.Task, error) {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	task.Title = title
	task.Completed = completed

	err = s.repo.Update(task)
	return task, err
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
