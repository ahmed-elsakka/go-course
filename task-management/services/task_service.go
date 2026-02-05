package services

import (
	"errors"
	"strings"
	"task-management/models"
	"task-management/repositories"
)

type TaskService struct {
	repo repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string) (*models.Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, errors.New("Error creating task: empty title")
	}

	task := &models.Task{
		Title:     title,
		Completed: false,
	}

	err := s.repo.Create(task)
	return task, err
}

func (s *TaskService) GetAll() ([]models.Task, error) {
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

func (s *TaskService) Delete(id uint) error {
	return s.repo.Delete(id)
}
