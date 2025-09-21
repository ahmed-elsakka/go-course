package main

import "fmt"

func increment(n *int) {
	*n = *n + 1
}

func main() {
	ptr := new(int)
	*ptr = 3
	increment(ptr)
	fmt.Println(*ptr)
}
