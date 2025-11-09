package main

import "fmt"

func main() {
	fmt.Println("Starting program")
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	panic("Something went terribly wrong!")
	fmt.Println("Ending program")
}
