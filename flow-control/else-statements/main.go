package main

import "fmt"

const (
	showName   string  = "Funny Show"
	appName    string  = "Ticket Booth"
	adultPrice float32 = 12.5453
	childPrice float32 = 8.5
	bookingFee float32 = 2.3
)

func multiply(a float32, b float32) float32 {
	return a * b
}

func displayAppName() {
	fmt.Println("===" + appName + "===")
}

func calcTotalPrice(adultCount int, childCount int) float32 {
	totalAdultPrice := multiply(float32(adultCount), adultPrice)
	totalChildrenPrice := multiply(float32(childCount), childPrice)
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

	if childCount+adultCount >= 1 {
		isGroup := adultCount+childCount >= 10 && adultCount >= 7
		totalPrice := calcTotalPrice(adultCount, childCount)
		fmt.Printf("Show Name: %s | Is Group Booking: %t | Total price is: %f", showName, isGroup, totalPrice)
	} else {
		fmt.Println("The number of people must be at least 1")
	}

}
