package main

import (
	"fmt"
	"os"
)

func main() {
	data, _ := os.ReadFile("example.txt")
	fmt.Println(string(data))
}
