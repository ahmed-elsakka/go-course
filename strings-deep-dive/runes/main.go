package main

import "fmt"

func main() {
	word := "Go语言" // "语言" means "language" in Chinese

	fmt.Println("Length in bytes: ", len(word))
	fmt.Println("Iterating byte by byte:")

	for i := 0; i < len(word); i++ {
		fmt.Printf("%c", word[i])
	}

	fmt.Println("\nIterating rune by rune:")
	for _, r := range word {
		fmt.Printf("%c", r)
	}
}
