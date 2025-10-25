package main

import "fmt"

func main() {
	var anyValue interface{}
	anyValue = 42
	anyValue = "Hello"
	anyValue = true

	fmt.Println(anyValue)

	var data []any // any is an alias to interface{}
	data = append(data, 10, "Go", true, 3.14)

	for _, v := range data {
		fmt.Println(v)
	}
}
