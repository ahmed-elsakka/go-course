package main

import (
	"fmt"
	"os"
)

func main() {
	writeAndReadFile()
	err := os.Rename("example.txt", "renamed_example.txt")
	if err != nil {
		fmt.Println("Error renaming file: ", err)
		return
	}
	fmt.Println("File renamed successfully!")

	err = os.Remove("renamed_example.txt")
	if err != nil {
		fmt.Println("Error deleting file: ", err)
		return
	}
	fmt.Println("File removed successfully!")
}

func writeAndReadFile() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("Hello World!")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}

	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	fmt.Println("File content:")
	fmt.Println(string(data))
}
