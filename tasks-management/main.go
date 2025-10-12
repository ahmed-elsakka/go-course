package main

import (
	"log"
	"net/http"
	"tasks-management/api"
	"tasks-management/daemons"
	"tasks-management/db"
	"tasks-management/models"
)

func main() {
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := &models.SQLRepo{DB: database}

	handler := &api.Handler{Repo: repo}
	http.HandleFunc("/tasks", handler.ListTasks)
	http.HandleFunc("/task/create", handler.CreateTask)

	daemons.StartNotifier()

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
