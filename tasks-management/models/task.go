package models

import "time"

type Task struct {
	ID          int
	UserID      int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type TaskRepository interface {
	Create(task *Task) error
	GetAll(userID int) ([]Task, error)
	Update(task *Task) error
	Delete(id int) error
}
