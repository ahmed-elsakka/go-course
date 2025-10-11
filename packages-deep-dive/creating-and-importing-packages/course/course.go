package course

import "fmt"

type Course struct {
	Code     string
	Title    string
	Enrolled int
}

func New(code, title string) Course {
	return Course{
		Code:  code,
		Title: title,
	}
}

func (c Course) Info() string {
	return fmt.Sprintf("%s %s - %d enrolled", c.Code, c.Title, c.Enrolled)
}
