package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	ID        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer response.Body.Close()

	var todo Todo
	if err := json.NewDecoder(response.Body).Decode(&todo); err != nil {
		fmt.Println("Error decoding json:", err)
		return
	}
	fmt.Printf("Todo id: %d, user id: %d, title: %s, completed: %t", todo.ID, todo.UserId, todo.Title, todo.Completed)
}
