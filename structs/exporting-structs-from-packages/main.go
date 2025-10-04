package main

import (
	"./student"
)

type Student struct {
	name string
	age  int
}

func main() {
	s := student.New("Sara", 21)
	s.PrintDetails()
}
