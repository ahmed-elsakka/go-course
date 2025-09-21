package main

import "fmt"

func main() {

	// Store scores in an array
	scores := [5]int{90, 85, 70, 88, 95}

	// Print the whole array
	fmt.Println("Scores:", scores)

	// Calculate average using a loop
	total := 0
	for _, val := range scores {
		total += val
	}
	avg := total / len(scores)

	fmt.Println("Average score:", avg)

}
