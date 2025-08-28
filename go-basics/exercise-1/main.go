package main

import "fmt"

const (
	appName    string = "Ticket Booth"
	showName   string = "Funny Show"
	adultPrice int    = 12
	childPrice int    = 8
	bookingFee int    = 2
)

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

func calcPriceAndDisplayInfo(adultCount int, childCount int) {
	totalAdultPrice := multiply(adultCount, adultPrice)
	totalChildrenPrice := multiply(childCount, childPrice)
	totalPrice := totalAdultPrice + totalChildrenPrice + bookingFee

	fmt.Printf("Show Name: %s | Total price is: %d", showName, totalPrice)
}

func main() {
	var (
		adultCount int
		childCount int
	)

	displayAppName()

	fmt.Println("Enter the number of adults:")
	fmt.Scanln(&adultCount)

	fmt.Println("Enter the number of children:")
	fmt.Scanln(&childCount)

	calcPriceAndDisplayInfo(adultCount, childCount)
}
