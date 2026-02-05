package main

import (
	"fmt"
	"net/http"
	"task-management/config"
	"task-management/handlers"
	"task-management/repositories"
	"task-management/services"
)

func main() {
	db, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println("Err: ", err)
	}

	repo := repositories.NewMySQLTaskRepository(db)
	service := services.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	http.HandleFunc("/tasks", handler.HandleTasks)
	http.HandleFunc("/tasks/", handler.HandleTaskById)

	http.ListenAndServe(":8000", nil)

}
