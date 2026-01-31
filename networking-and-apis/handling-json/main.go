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
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error making api request: ", err)
		return
	}

	defer res.Body.Close()

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding json response: ", err)
		return
	}
	fmt.Printf("Todo id: %d, user id: %d, title: %s, completed: %t", todo.ID, todo.UserId, todo.Title, todo.Completed)
}
