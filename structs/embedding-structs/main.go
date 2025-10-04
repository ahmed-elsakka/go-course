package main

import "fmt"

type Address struct {
	street  string
	city    string
	country string
}

type User struct {
	name string
	age  int
	Address
}

func main() {
	user := User{
		name: "Sara",
		age:  21,
		Address: Address{
			street:  "123 Main St",
			city:    "London",
			country: "UK",
		},
	}
	fmt.Println("User: ", user.name)
	fmt.Println("Age: ", user.age)
	fmt.Printf("Address: %s, %s, %s", user.street, user.city, user.country)
}
