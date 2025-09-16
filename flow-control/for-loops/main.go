package main

import "fmt"

func main() {
	var enteredPin int
	const userPin = 1234

	for attempt := 1; attempt <= 3; attempt++ {
		fmt.Println("Enter  the PIN: ")
		fmt.Scanln(&enteredPin)

		if enteredPin == userPin {

		} else {
			fmt.Println("Incorrect PIN")
		}
	}
}
