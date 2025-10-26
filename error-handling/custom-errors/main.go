package main

import (
	"fmt"
)

type ValidationError struct {
	Field string
	Msg   string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field '%s': %s", v.Field, v.Msg)
}

func validateUser(name string, age int) error {
	if name == "" {
		return ValidationError{Field: "name", Msg: "cannot be empty"}
	}
	if age <= 0 {
		return ValidationError{Field: "age", Msg: "must be positive"}
	}
	return nil
}

func main() {
	err := validateUser("", 25)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User is valid")
}
