package main

import "fmt"

func main() {
	fmt.Println("Start")

	defer fmt.Println("Deferred call 1")
	defer fmt.Println("Deferred call 2")

	fmt.Println("End")
}
