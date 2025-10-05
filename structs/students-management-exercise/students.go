package main

import "fmt"

type Person struct {
	Name  string
	Email string
}

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
	Person
}

type Course struct {
	Code     string
	Title    string
	Enrolled int
}

func NewStudent(id, name, email string) Student {
	return Student{
		ID:              id,
		Status:          StatusApplicant,
		NumberOfCourses: 0,
		Person: Person{
			Name:  name,
			Email: email,
		},
	}
}

func NewCourse(code, title string) Course {
	return Course{
		Code:  code,
		Title: title,
	}
}

func (s Student) Summary() string {
	return fmt.Sprintf("%s [%s] - %s", s.Name, s.Email, s.Status)
}

func (c Course) Info() string {
	return fmt.Sprintf("%s %s - %d enrolled", c.Code, c.Title, c.Enrolled)
}

func (s *Student) Activate() {
	s.Status = StatusActive
}

func (s *Student) Enroll(c *Course) {
	c.Enrolled++
	s.NumberOfCourses++
}

func (s *Student) Drop(c *Course) {
	c.Enrolled--
	s.NumberOfCourses--
}
