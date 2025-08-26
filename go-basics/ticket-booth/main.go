package main

import "fmt"

const (
	currency   = "£"
	adultPrice = 12
	childPrice = 8
	bookingFee = 1
	showName   = "Funny Show"
	appName    = "Ticket Booth"
)

func main() {
	var adultCount int
	var childCount int

	fmt.Print("Enter number of adults: ")
	fmt.Scanln(&adultCount)

	fmt.Print("Enter number of children: ")
	fmt.Scanln(&childCount)

	isEveningShow := true
	isGroup := (adultCount + childCount) >= 3

	totalAdults := adultCount * adultPrice
	totalChildren := childCount * childPrice
	subtotal := totalAdults + totalChildren
	total := subtotal + bookingFee

	fmt.Println("===" + appName + "===")
	fmt.Println("\n*** Show Information ***")
	fmt.Printf("Show Name: %s | Evening Show: %t  |  Group(>=3): %t\n", showName, isEveningShow, isGroup)

	fmt.Println("\n*** Ticket Information ***")
	summary := fmt.Sprintf("%s — tickets %d (adults %d, children %d) → Total %s%d",
		showName, adultCount+childCount, adultCount, childCount, currency, total)
	fmt.Println(summary)
}
