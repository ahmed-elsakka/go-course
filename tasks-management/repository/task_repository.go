package repository

import "task-manager/models"

type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	GetByID(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
}
