package models

import "testing"

func TestCreateTask(t *testing.T) {
	repo := NewInMemoryRepo()
	task := &Task{Title: "Test", Description: "Demo"}
	err := repo.Create(task)
	if err != nil {
		t.Fatal("Expected nil, got", err)
	}
	if len(repo.tasks) != 1 {
		t.Error("Expected 1 task, got", len(repo.tasks))
	}
}
