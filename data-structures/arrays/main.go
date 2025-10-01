package main

import "fmt"

func main() {
	/* Creating arrays */
	//var numbers [5]int
	//var numbers = [5]int{1, 2, 3, 4, 5}
	numbers := [...]int{1, 2, 3, 4, 5}

	/* Accessing arrays elements */
	//fmt.Println(numbers[2])
	numbers[1]++
	// fmt.Println(numbers)

	/* Iterating arrays elements */
	/* Tranditional loop */
	for i := 0; i < len(numbers); i++ {
		//fmt.Printf("%d  ", numbers[i])
	}

	/* Using range */
	for _, val := range numbers {
		//fmt.Printf("Index: %d, value: %d  \n", index, val)
		fmt.Printf("%d ", val)
	}

}
