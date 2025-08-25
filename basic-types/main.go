package main

import "fmt"

const (
	Standard = iota
	VIP
	Premium
)

const (
	appName    = "Ticket Booth"
	adultPrice = 12.50
	childPrice = 8.25
	bookingFee = 1.00
)

func printFormattedLine(label string, amount float64) {
	fmt.Printf("%-12s £%.2f\n", label+":", amount)
}

func getCost(unitPrice float64, count int) float64 {
	return float64(count) * unitPrice
}

func main() {
	fmt.Println("===", appName, "===")
	var (
		showName   string
		adultCount int
		childCount int
		seatType   int = Standard
	)

	fmt.Print("Show name:")
	fmt.Scanln(&showName)

	fmt.Print("Adults: ")
	fmt.Scanln(&adultCount)

	fmt.Print("Children: ")
	fmt.Scanln(&childCount)

	isMorningShow := true
	isGroup := (adultCount + childCount) >= 3
	isVip := (seatType == VIP)

	totalAdultsCost := getCost(adultPrice, adultCount)
	totalChildrenCost := getCost(childPrice, childCount)

	totalCost := totalAdultsCost + totalChildrenCost

	fmt.Println()
	fmt.Printf("Seat type code: %d (Standard=0, VIP=1, Premium=2)\n", seatType)
	fmt.Printf("Is morning show: %t | is group: %t | is VIP: %t\n", isMorningShow, isGroup, isVip)
	fmt.Printf("%s - %d adults, %d children\n\n", showName, adultCount, childCount)
	printFormattedLine("Adults", totalAdultsCost)
	printFormattedLine("Children", totalChildrenCost)
	printFormattedLine("Booking fee", bookingFee)
	printFormattedLine("Total", totalCost)

}
