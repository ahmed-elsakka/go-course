package main

import "fmt"

func main() {
	students := map[int]string{
		1: "Sara",
		2: "John",
		3: "Jessica",
	}

	students[4] = "Bob"
	students[2] = "Alice"

	delete(students, 3)
	name, exists := students[1]
	if exists {
		fmt.Println("Found: ", name)
	} else {
		fmt.Println("Not found")
	}
}
