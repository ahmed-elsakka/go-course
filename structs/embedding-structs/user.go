package main

import (
	"errors"
	"fmt"
)

type User struct {
	name  string
	age   int
	email string
	Address
}

type Address struct {
	street  string
	city    string
	country string
}

func NewUser(name string, age int, email string, address Address) (*User, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if age <= 0 {
		return nil, errors.New("age is not correct")
	}
	if email == "" || !validateEmail(email) {
		return nil, errors.New("email is not correct")
	}
	return &User{name: name, age: age, email: email, Address: address}, nil
}

func (user *User) process() {
	fmt.Printf("Name: %s, age: %d, email: %s\n", user.name, user.age, user.email)
	fmt.Printf("Street: %s, city: %s, country: %s", user.Address.street, user.Address.city, user.Address.country)
}

func (user *User) updateEmail(newEmail string) error {
	if !validateEmail(newEmail) {
		return errors.New("invalid email")
	}
	user.email = newEmail
	return nil
}

func validateEmail(email string) bool {
	for _, c := range email {
		if c == '@' {
			return true
		}
	}
	return false
}
