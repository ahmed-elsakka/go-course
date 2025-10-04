package student

import "fmt"

type Student struct {
	name string
	age  int
}

func New(name string, age int) Student {
	return Student{name: name, age: age}
}

func (student Student) PrintDetails() {
	fmt.Printf("Name: %s , age: %d\n", student.name, student.age)
}

func (student *Student) SetAge(newAge int) {
	student.age = newAge
}
