package main

import "fmt"

// Go Basics: Exercise 1
// 1. Create a function to calculate the total tickets price
// 2. Display the show name and the total price instead of displaying the total price only

const (
	showName   string = "Funny Show"
	appName    string = "Ticket Booth"
	adultPrice int    = 12
	childPrice int    = 8
	bookingFee int    = 2
)

func multiply(a int, b int) int {
	return a * b
}

func displayAppName() {
	fmt.Println("===" + appName + "===")
}

func calcTotalPrice(adultCount int, childCount int) int {
	totalAdultPrice := multiply(adultCount, adultPrice)
	totalChildrenPrice := multiply(childCount, childPrice)
	return totalAdultPrice + totalChildrenPrice + bookingFee
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

	totalPrice := calcTotalPrice(adultCount, childCount)
	fmt.Printf("Show Name: %s | Total price is: %d", showName, totalPrice)
}
