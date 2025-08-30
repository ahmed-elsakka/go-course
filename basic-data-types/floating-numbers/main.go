package main

import "fmt"

const (
	showName   string  = "Funny Show"
	appName    string  = "Ticket Booth"
	adultPrice float32 = 12.5
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

	totalPrice := calcTotalPrice(adultCount, childCount)
	fmt.Printf("Show Name: %s | Total price is: %f", showName, totalPrice)
}
