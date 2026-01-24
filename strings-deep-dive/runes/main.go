package main

import "fmt"

func main() {
	str := "Hey💻"
	runes := []rune(str)
	runes[1] = 't'
	var r rune = runes[1]
	fmt.Printf("%c", r)
}
