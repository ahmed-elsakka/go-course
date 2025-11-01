package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // ensures file is closed even if later code fails

	fmt.Println("File opened successfully")

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		// simulate a bug or unexpected condition
		if line == "STOP" {
			fmt.Println("Encountered STOP line — aborting early.")
			return // <- file will STILL be closed here because of defer!
		}

		fmt.Printf("Line %d: %s\n", lineCount, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("File read complete — %d lines processed.\n", lineCount)
}

/*func main() {
	fmt.Println("Start")

	defer fmt.Println("Deferred call 1")

	fmt.Println("End")
}*/
