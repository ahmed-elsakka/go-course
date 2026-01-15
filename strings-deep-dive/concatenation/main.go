package main

import (
	"fmt"
	"strings"
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

	age := 25

	message = fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(message)

	var sb strings.Builder

	for i := 1; i <= 5; i++ {
		sb.WriteString(fmt.Sprintf("Line %d\n", i))
	}

	fmt.Println(sb.String())

}
