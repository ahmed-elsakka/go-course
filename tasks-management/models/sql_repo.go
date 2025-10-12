package models

import (
	"database/sql"
	"time"
)

type SQLRepo struct {
	DB *sql.DB
}

func (r *SQLRepo) Create(task *Task) error {
	_, err := r.DB.Exec(
		"INSERT INTO tasks (user_id, title, description, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		task.UserID, task.Title, task.Description, task.Status, time.Now(), time.Now(),
	)
	return err
}

func (r *SQLRepo) GetAll(userID int) ([]Task, error) {
	rows, err := r.DB.Query("SELECT id, user_id, title, description, status, created_at, updated_at FROM tasks WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.ID, &t.UserID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *SQLRepo) Update(task *Task) error {
	_, err := r.DB.Exec(
		"UPDATE tasks SET title=?, description=?, status=?, updated_at=? WHERE id=?",
		task.Title, task.Description, task.Status, time.Now(), task.ID,
	)
	return err
}

func (r *SQLRepo) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
