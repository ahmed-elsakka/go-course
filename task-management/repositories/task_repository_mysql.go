package repositories

import (
	"task-management/models"

	"gorm.io/gorm"
)

type MySQLTaskRepository struct {
	db *gorm.DB
}

func NewMySQLTaskRepository(db *gorm.DB) TaskRepository {
	return &MySQLTaskRepository{db: db}
}

func (r *MySQLTaskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *MySQLTaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *MySQLTaskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	res := r.db.First(&task, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &task, res.Error
}

func (r *MySQLTaskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *MySQLTaskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}
