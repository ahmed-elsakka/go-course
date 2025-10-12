package main

import (
	"fmt"
	"os"
)

func safeFileRead(filename string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	file, err := os.Open(filename)
	if err != nil {
		panic("Failed to open file: " + err.Error())
	}
	defer file.Close()

	fmt.Println("File opened successfully!")
}

func main() {
	safeFileRead("nonexistent.txt")
	fmt.Println("Program completed safely.")
}
