package main

import "fmt"

func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(Max(10, 20))
	fmt.Println(Max(3.5, 2.7))
	fmt.Println(Max("Go", "Golang"))
}
