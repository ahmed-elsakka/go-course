package main

import (
	"fmt"
	"students-system/course"
	"students-system/student"
)

func main() {
	fmt.Println("Student Management Demo")

	student := student.New("s-1", "Alice", "alice@email.com")
	mathCourse := course.New("MATH101", "Math")

	fmt.Println(student.Summary())

	student.Activate()
	fmt.Println("Student activated: ", student.Summary())

	student.Enroll(&mathCourse)
	fmt.Println("Course info:", mathCourse.Info())

}
