package main

import "fmt"

type User struct {
	name  string
	age   int
	email string
}

func (user *User) process() {
	fmt.Printf("Name: %s, age: %d, email: %s", user.name, user.age, user.email)
}

func (user *User) updateEmail(newEmail string) {
	user.email = newEmail
}
