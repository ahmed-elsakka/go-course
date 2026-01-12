package main

import "fmt"

func main() {
	str1 := "Hello world"
	// str1[0] = 'h'
	fmt.Println(str1)

	str2 := `Hello

	World
	`
	fmt.Println(str2)

	fmt.Println(len(str1))
	fmt.Println(len(str2))
}
