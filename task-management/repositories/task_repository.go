package repositories

import "task-management/models"

type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
	GetByID(id uint) (*models.Task, error)
}
