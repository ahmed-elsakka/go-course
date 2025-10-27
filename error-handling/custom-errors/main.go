package main

import (
	"fmt"
	"strings"
)

type UserValidationError struct {
	Field string
}

func (err UserValidationError) Error() string {
	return fmt.Sprintf("invalid user %s", err.Field)
}

func validateUserData(name string, email string) error {
	if name == "" {
		return UserValidationError{Field: "name"}
	}
	if !strings.Contains(email, "@") {
		return UserValidationError{Field: "email"}
	}
	return nil
}

func registerUser(name string, email string) {

	if err := validateUserData(name, email); err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("User %s (email %s) registered successfully!\n", name, email)
}

func main() {
	registerUser("", "incorrectEmail")
	registerUser("Alice", "alice@example.com")
	registerUser("Sara", "sara.com")
}
