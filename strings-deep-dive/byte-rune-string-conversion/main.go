package main

import "fmt"

func main() {
	word := "Go语言" // "语言" means "language" in Chinese

	bytes := []byte(word)
	fmt.Println(bytes)

	runes := []rune(word)
	//fmt.Println(runes)
	fmt.Printf("%U", runes)
}
