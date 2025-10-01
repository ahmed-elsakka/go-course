package main

import "fmt"

func main() {
	arr := [6]int{10, 20, 30, 40, 50, 60}
	var slice []int = arr[1:4]
	fmt.Println(slice)
	slice[0] = 43
	arr[2] = 32
	fmt.Println("Modified Slice:", slice)
	fmt.Println("Modified Array:", arr)

	// numsSlice := []int{1, 2, 3}

	numsSlice := make([]int, 3, 5)
	numsSlice[2] = 22
	fmt.Println("numsSlice:", numsSlice)
	fmt.Println("Length:", len(numsSlice))
	fmt.Println("Capacity:", cap(numsSlice))

}
