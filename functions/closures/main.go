package main

import "fmt"

func getAmountFromUser(balance float64, action string) (amount float64, err error) {
	fmt.Printf("Enter the amount: ")
	fmt.Scanln(&amount)
	if amountValidator(amount, balance, action) {
		err = nil
		return
	}
	err = fmt.Errorf("Amount is not valid.")
	return
}

var amountValidator = func(amount float64, balance float64, action string) bool {
	if amount <= 0 {
		fmt.Println("Wrong amount")
		return false
	}
	if amount > balance && action == "withdraw" {
		fmt.Println("Insufficient balance")
		return false
	}
	return true
}

func balanceManager(initialAmount float64) func(action string) float64 {
	balance := initialAmount
	return func(action string) float64 {
		if action == "retrieve" {
			return balance
		}
		amount, err := getAmountFromUser(balance, action)
		if err != nil {
			return balance
		}
		if action == "withdraw" {
			balance -= amount
			fmt.Printf("Deposit is done")
		} else if action == "deposit" {
			balance += amount
			fmt.Printf("Withdrawal is done")
		}
		return balance
	}
}

func main() {
	var enteredPin int
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

	var manageBalance = balanceManager(1000)
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
			fmt.Printf("Balance: %.2f", manageBalance("retrieve"))
		case 2:
			manageBalance("deposit")
		case 3:
			manageBalance("withdraw")
		case 4:
			break outerLoop
		default:
			fmt.Println("Incorrect choice")
		}
	}

}
