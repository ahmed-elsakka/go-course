package main

import (
	"fmt"
)

func main() {
	greeting := "Hello"
	name := "Alice"

	message := greeting + ", " + name + "!"
	fmt.Println(message)

	message = "Hello"
	message += " "
	message += "World"
	fmt.Println(message)
}
