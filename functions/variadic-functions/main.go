package main

import "fmt"

func addAll(numbers ...int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func main() {
	sum := addAll()
	fmt.Println(sum)
}
