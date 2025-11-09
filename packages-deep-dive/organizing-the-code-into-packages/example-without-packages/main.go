package main

import (
	"fmt"
	"students-system/course"
	"students-system/student"
)

func main() {
	fmt.Println("Student Management Demo")

	s := student.NewStudent("s-1", "Alice", "alice@email.com")
	mathCourse := course.NewCourse("MATH101", "Math")

	fmt.Println(s.Summary())

	s.Activate()
	fmt.Println("Student activated: ", s.Summary())

	s.Enroll(&mathCourse)
	fmt.Println("Course info:", mathCourse.Info())

}
