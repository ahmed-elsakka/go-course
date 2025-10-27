package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateUserData(name string, email string) error {
	if name == "" {
		return errors.New("invalid user name")
	}
	if !strings.Contains(email, "@") {
		return fmt.Errorf("invalid email: %s", email)
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
	registerUser("", "email@example.com")
}
