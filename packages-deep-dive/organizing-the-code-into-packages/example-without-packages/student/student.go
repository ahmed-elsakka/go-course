package student

import (
	"fmt"
	"students-system/course"
	"students-system/person"
)

const (
	StatusApplicant = "applicant"
	StatusActive    = "active"
	StatusAlumni    = "alumni"
)

type Student struct {
	ID              string
	Status          string
	GPA             float64
	NumberOfCourses int
	person.Person
}

func (s Student) Summary() string {
	return fmt.Sprintf("%s [%s] - %s", s.Name, s.Email, s.Status)
}

func (s *Student) Activate() {
	s.Status = StatusActive
}

func (s *Student) Enroll(c *course.Course) {
	c.Enrolled++
	s.NumberOfCourses++
}

func (s *Student) Drop(c *course.Course) {
	c.Enrolled--
	s.NumberOfCourses--
}

func NewStudent(id, name, email string) Student {
	return Student{
		ID:              id,
		Status:          StatusApplicant,
		NumberOfCourses: 0,
		Person: person.Person{
			Name:  name,
			Email: email,
		},
	}
}
