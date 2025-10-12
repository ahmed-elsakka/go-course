package main

import "fmt"

func add(a, b int) int {
	sum := a + b
	return sum
}

func main() {
	x, y := 10, 20
	result := add(x, y)
	fmt.Println("Sum:", result)
}
