package main

import "fmt"

const appName string = "Ticket Booth"

// Function definition
func multiply(a int, b int) int {
	return a * b
}

/*
This function displays the app name
The appName constant contains the application name
*/
func displayAppName() {
	fmt.Println("===" + appName + "===")
}

func main() {
	const (
		showName   string = "Funny Show"
		adultPrice int    = 12
		childPrice int    = 8
		bookingFee int    = 2
	)

	var (
		adultCount int
		childCount int
	)

	fmt.Println("===" + appName + "===")
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
