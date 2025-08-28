package main

import "fmt"

// Function definition
func multiply(a int, b int) int {
	return a * b
}

func displayAppName() {
	fmt.Println("=== Ticket Booth ===")
}

func main() {
	const (
		appName    string = "Ticket Booth"
		showName   string = "Funny Show"
		adultPrice int    = 12
		childPrice int    = 8
		bookingFee int    = 2
	)

	var (
		adultCount int
		childCount int
	)

	displayAppName()

	fmt.Println("Enter the number of adults:")
	fmt.Scanln(&adultCount)

	fmt.Println("Enter the number of children:")
	fmt.Scanln(&childCount)

	totalAdultPrice := multiply(adultCount, adultPrice)
	totalChildrenPrice := multiply(childCount, childPrice)
	totalPrice := totalAdultPrice + totalChildrenPrice + bookingFee

	fmt.Printf("Total price is: %d", totalPrice)
}
