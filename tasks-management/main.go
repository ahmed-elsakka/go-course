package main

import (
	"net/http"
	"task-manager/config"
	"task-manager/handlers"
	"task-manager/models"
	"task-manager/repository"
	"task-manager/service"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Task{})

	repo := repository.NewMySQLTaskRepository(db)
	service := service.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	http.HandleFunc("/tasks", handler.HandleTasks)
	http.HandleFunc("/tasks/", handler.HandleTaskByID)

	http.ListenAndServe(":8080", nil)
}
