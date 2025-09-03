package main

import "fmt"

func main() {
	var (
		a float64
		b float64
	)

	fmt.Println("Enter the first number: ")
	fmt.Scanln(&a)

	fmt.Println("Enter the second number: ")
	fmt.Scanln(&b)

	fmt.Printf("The sum of the two numbers is: %.2f\n", a+b)
	fmt.Printf("The first number is larger than the second: %t", a > b)
}
