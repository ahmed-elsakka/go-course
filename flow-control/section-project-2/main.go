package main

import (
	"fmt"
)

func main() {
	const pin = 1234
	var enteredPin int

Login: // Label for login block
	attempts := 0

	for attempts < 3 {
		fmt.Print("Enter PIN: ")
		fmt.Scan(&enteredPin)

		if enteredPin == pin {
			fmt.Println("✅ Login successful!")
			break
		} else {
			fmt.Println("❌ Incorrect PIN")
		}
		attempts++
	}

	if attempts == 3 {
		fmt.Println("🚨 Too many attempts. Account locked.")
		fmt.Println("Do you want to retry? (1=yes, 2=no): ")
		var retry int
		fmt.Scan(&retry)
		if retry == 1 {
			goto Login // Restart login process
		} else {
			return
		}
	}

	// Infinite loop for ATM menu
	balance := 1000.0
MenuLoop: // Label for the ATM menu loop
	for {
		fmt.Println("\n--- ATM Menu ---")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Choose option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("💰 Balance: $%.2f\n", balance)
		case 2:
			var amount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("⚠️ Deposit must be greater than 0.")
				continue
			}
			balance += amount
			fmt.Println("✅ Deposit successful.")
		case 3:
			var amount float64
			fmt.Print("Enter withdrawal amount: ")
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("⚠️ Withdrawal must be greater than 0.")
				continue
			}
			if amount > balance {
				fmt.Println("⚠️ Insufficient funds.")
				continue
			}
			balance -= amount
			fmt.Println("✅ Withdrawal successful.")
		case 4:
			fmt.Println("👋 Thank you. Goodbye!")
			break MenuLoop
		default:
			fmt.Println("❌ Invalid option.")
		}
	}
}
