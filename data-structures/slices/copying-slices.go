package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	dst := make([]int, 5)
	n := copy(dst, src)

	fmt.Println("Source:", src)
	fmt.Println("Dest:", dst)
	fmt.Println("Number of elements copied:", n)
}
