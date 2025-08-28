package main

import "fmt"

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
	fmt.Println("=== ", appName, " ===")

	fmt.Println("Enter the number of adults:")
	fmt.Scanln(&adultCount)

	fmt.Println("Enter the number of children:")
	fmt.Scanln(&childCount)

	totalAdultPrice := adultCount * adultPrice
	totalChildrenPrice := childCount * childPrice
	totalPrice := totalAdultPrice + totalChildrenPrice + bookingFee

	fmt.Printf("Total price is: %d", totalPrice)
}
