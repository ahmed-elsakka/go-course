package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("Index %d → Value %d\n", i, nums[i])
	}
	fmt.Println("")
	for i, v := range nums {
		fmt.Printf("Index %d → Value %d\n", i, v)
	}
}
