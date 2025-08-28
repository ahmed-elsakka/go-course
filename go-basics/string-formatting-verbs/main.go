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
		adultCount    int  = 3
		childCount    int  = 2
		isEveningShow bool = false
	)

	totalAdultPrice := adultCount * adultPrice
	totalChildrenPrice := childCount * childPrice
	totalPrice := totalAdultPrice + totalChildrenPrice + bookingFee

	fmt.Println("=== ", appName, " ===")
	showInfo := fmt.Sprintf("Show name: %s | Is evening show: %t", showName, isEveningShow)
	fmt.Println(showInfo)
	fmt.Printf("Total price is: %d", totalPrice)
}
