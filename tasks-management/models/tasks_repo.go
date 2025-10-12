package models

import "errors"

type InMemoryRepo struct {
	tasks []Task
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{tasks: []Task{}}
}

func (r *InMemoryRepo) Create(task *Task) error {
	task.ID = len(r.tasks) + 1
	r.tasks = append(r.tasks, *task)
	return nil
}

func (r *InMemoryRepo) GetAll(userID int) ([]Task, error) {
	return r.tasks, nil
}

func (r *InMemoryRepo) Update(task *Task) error {
	for i := range r.tasks {
		if r.tasks[i].ID == task.ID {
			r.tasks[i] = *task
			return nil
		}
	}
	return errors.New("task not found")
}

func (r *InMemoryRepo) Delete(id int) error {
	for i := range r.tasks {
		if r.tasks[i].ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
