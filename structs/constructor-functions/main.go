package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	s := newStudent("Sara", 21)
	s.printStudentDetails()
}

func newStudent(name string, age int) Student {
	return Student{name: name, age: age}
}

func (student Student) printStudentDetails() {
	fmt.Printf("Name: %s , age: %d\n", student.name, student.age)
}

func (student *Student) setAge(newAge int) {
	student.age = newAge
}
