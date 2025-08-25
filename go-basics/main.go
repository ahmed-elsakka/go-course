package main

import "fmt"

const (
	appName    = "Ticket Booth"
	adultPrice = 12
	childPrice = 8
	bookingFee = 1
)

func printFormattedLine(label string, amount int) {
	fmt.Printf("%-12s £%d\n", label+":", amount)
}

func getCost(unitPrice int, count int) int {
	return count * unitPrice
}

func main() {
	fmt.Println("===", appName, "===")
	var (
		showName   string
		adultCount int
		childCount int
	)

	fmt.Print("Show name:")
	fmt.Scanln(&showName)

	fmt.Print("Adults: ")
	fmt.Scanln(&adultCount)

	fmt.Print("Children: ")
	fmt.Scanln(&childCount)

	totalAdultsCost := getCost(adultCount, adultPrice)
	totalChildrenCost := getCost(childCount, childPrice)

	totalCost := totalAdultsCost + totalChildrenCost

	fmt.Println()
	fmt.Printf("%s - %d adults, %d children\n\n", showName, adultCount, childCount)
	printFormattedLine("Adults", totalAdultsCost)
	printFormattedLine("Children", totalChildrenCost)
	printFormattedLine("Booking fee", bookingFee)
	printFormattedLine("Total", totalCost)

}
