package main

import (
	"fmt"
)

type User struct {
	name  string
	age   int
	email string
}

func main() {
	user := User{name: "Sara", age: 21, email: "sara@email.com"}
	updateUserEmail(&user)
	processUser(&user)
}

func processUser(user *User) {
	fmt.Printf("Name: %s, age: %d, email: %s", user.name, user.age, user.email)
}

func updateUserEmail(user *User) {
	user.email = "example@email.com"
}
