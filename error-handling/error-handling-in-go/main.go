package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:")
	fmt.Println(string(data))
}
