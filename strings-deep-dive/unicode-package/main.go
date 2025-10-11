package main

import (
	"fmt"
	"unicode"
)

func main() {
	text := "Go语123🙂"

	for _, r := range text {
		switch {
		case unicode.IsLetter(r):
			fmt.Printf("%q is a letter.\n", r)
		case unicode.IsDigit(r):
			fmt.Printf("%q is a digit.\n", r)
		case unicode.IsSpace(r):
			fmt.Printf("%q is a space.\n", r)
		default:
			fmt.Printf("%q is a symbol or other character.\n", r)
		}
	}
}
