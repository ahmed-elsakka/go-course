package main

import "fmt"

func getAmountFromUser(balance float64) (amount float64, err error) {
	fmt.Printf("Enter the amount: ")
	fmt.Scanln(&amount)
	if amountValidator(amount, balance) {
		err = nil
		return
	}
	err = fmt.Errorf("Amount is not valid.")
	return
}

var amountValidator = func(amount float64, balance float64) bool {
	if amount <= 0 {
		fmt.Println("Wrong amount")
		return false
	}
	if amount > balance {
		fmt.Println("Insufficient balance")
		return false
	}
	return true
}

func main() {
	var (
		enteredPin int
		balance    float64 = 1000.0
	)
	const userPin = 1234

login:
	attempt := 1

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

	if attempt == 4 {
		var retry int
		fmt.Println("Attempts to enter PIN exceeded the maximum, retry (1. yes, 2. no)")
		fmt.Scanln(&retry)
		if retry == 1 {
			goto login
		} else {
			return
		}
	}

outerLoop:
	for {
		fmt.Println("\n--- ATM Menu ---")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Printf("Balance: %.2f", balance)
		case 2:
			amount, err := getAmountFromUser(balance)
			if err == nil {
				balance += amount
				fmt.Printf("Deposit is done")
			}
		case 3:
			amount, err := getAmountFromUser(balance)
			if err == nil {
				balance -= amount
				fmt.Printf("Withdrawal is done")
			}
		case 4:
			break outerLoop
		default:
			fmt.Println("Incorrect choice")
		}
	}

}
