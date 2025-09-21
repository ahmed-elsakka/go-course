package main

import "fmt"

func main() {
	var ptr *int
	x := 41
	ptr = &x
	*ptr = 2
	fmt.Println(x)
}
