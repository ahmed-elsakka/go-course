package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func main() {
	newPost := Post{
		Title:  "Learning Go APIs",
		Body:   "This is an example post using Go!",
		UserID: 1,
	}

	postData, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts",
		"application/json", bytes.NewBuffer(postData))
	if err != nil {
		fmt.Println("Error sending POST:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println("Response from API:", result)
}
