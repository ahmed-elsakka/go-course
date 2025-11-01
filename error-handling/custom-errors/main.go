package main

import (
	"fmt"
	"strings"
)

type UserDataValidationError struct {
	Field string
}

func (e UserDataValidationError) Error() string {
	return fmt.Sprintf("invalid %s", e.Field)
}

func validateUserData(name string, email string) error {
	if name == "" {
		return UserDataValidationError{Field: "name"}
	}
	if !strings.Contains(email, "@") {
		return UserDataValidationError{Field: "email"}
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
	registerUser("Sara", "emailexample.com")
}
