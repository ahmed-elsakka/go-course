package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	name := "Sara"
	age := 21

	printStudentDetails(name, age)
}

func printStudentDetails(name string, age int) {
	fmt.Printf("Name: %s , age: %d", name, age)
}
