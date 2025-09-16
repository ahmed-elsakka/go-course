package main

import "fmt"

func main() {
	var enteredPin int
	const userPin = 1234
	attempt := 1

	// while a condition is true, keep executing
	for attempt <= 3 {
		fmt.Println("Enter the PIN: ")
		fmt.Scanln(&enteredPin)

		if enteredPin == userPin {
			fmt.Println("Logged in")
			break
		} else {
			fmt.Println("Incorrect PIN")
		}

		attempt++
	}
}
