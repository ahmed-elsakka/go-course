package main

import "fmt"

func increment(n *int) {
	*n = *n + 1
}

func main() {
	x := 5
	increment(&x)
	fmt.Println(x)
}
