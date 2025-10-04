package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 2
	b := 3

	swap(&a, &b)

	fmt.Printf("a is: %d , b is : %d\n", a, b)
}
