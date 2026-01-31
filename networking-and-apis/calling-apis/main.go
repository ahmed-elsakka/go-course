package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error making api request: ", err)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}
	fmt.Println(string(body))

}
