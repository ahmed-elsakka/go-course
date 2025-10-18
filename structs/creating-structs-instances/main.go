package main

import "fmt"

type User struct {
	name string
	age  int
}

func main() {
	user := User{}
	user.age = 21
	user.name = "Sara"

	user2 := new(User)
	user2.age = 21
	user2.name = "Bob"
	fmt.Printf("Name: %s, age: %d", user2.name, user2.age)
}
