package main

import "fmt"

func main() {
	var anyValue any
	anyValue = 12
	anyValue = "Test"
	anyValue = true
	fmt.Println(anyValue)

	var data []any
	data = append(data, 10, "test", true)
	fmt.Println(data)
}
