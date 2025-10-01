package main

import "fmt"

func main() {
	var (
		numOfStudents int
		average       float64
		highest       int
	)
	fmt.Println("Enter the number of students: ")
	fmt.Scanln(&numOfStudents)
	grades := make([]int, numOfStudents)
	for i := 0; i < numOfStudents; i++ {
		fmt.Printf("Enter the grade of the %dth student: \n", i+1)
		fmt.Scanln(&grades[i])
	}
	average = calcAverage(grades)
	highest = calcHeighest(grades)
	fmt.Println("The highest grade is: ", highest)
	fmt.Println("The average grade is: ", average)
}

func calcAverage(grades []int) float64 {
	var total int
	for _, val := range grades {
		total += val
	}
	return float64(total) / float64(len(grades))
}

func calcHeighest(grades []int) int {
	var highest int = 0
	for _, val := range grades {
		if val > highest {
			highest = val
		}
	}
	return highest
}
