package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateUser(name string, email string) error {

	if name == "" {
		return errors.New("name cannot be empty")
	}

	if !strings.Contains(email, "@") {
		return fmt.Errorf("invalid email: %s", email)
	}

	return nil
}

func registerUser(name string, email string) {
	err := validateUser(name, email)
	if err != nil {
		fmt.Println("Registration failed:", err)
		return
	}

	fmt.Printf("User %s (email %s) registered successfully!\n", name, email)
}

func main() {
	registerUser("", "incorrectEmail")
	registerUser("Alice", "aliceexample.com")
	registerUser("", "email@example.com")
}
