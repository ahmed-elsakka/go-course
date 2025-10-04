package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {

	s1 := Student{"Sara", 21}             // positional
	s2 := Student{name: "Alice", age: 23} // named fields
	var s3 Student                        // zero values
	s3.name = "Bob"
	s3.age = 30

	s1.setAge(22)
	s1.printStudentDetails()
	s2.printStudentDetails()
	s3.printStudentDetails()
}

func (student Student) printStudentDetails() {
	fmt.Printf("Name: %s , age: %d\n", student.name, student.age)
}

func (student *Student) setAge(newAge int) {
	student.age = newAge
}
