package main

import "fmt"

func main() {
	students := map[int]string{
		1: "Sara",
		2: "John",
		3: "Jessica",
	}

	// read entries
	name := students[5]
	fmt.Println(name)
	// add/update entries
	students[4] = "Alice"
	students[2] = "Bob"
	// delete entries
	delete(students, 3)

	fmt.Println(students)
}
