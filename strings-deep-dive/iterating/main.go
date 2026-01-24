package main

import "fmt"

func main() {
	str := "Go💻"

	for _, r := range str {
		fmt.Printf("%c", r)
	}

}
