package main

import "fmt"

func main() {
	fmt.Println("Student Management Demo")

	student := NewStudent("s-1", "Alice", "alice@email.com")
	mathCourse := NewCourse("MATH101", "Math")

	fmt.Println(student.Summary())

	student.Activate()
	fmt.Println("Student activated: ", student.Summary())

	student.Enroll(&mathCourse)
	fmt.Println("Course info:", mathCourse.Info())

}
