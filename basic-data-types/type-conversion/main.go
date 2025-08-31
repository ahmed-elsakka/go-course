package main

import "fmt"

func main() {
	/*var f float32 = 12.53
	var i int = int(f)
	fmt.Printf("Value of i is %d", i)*/

	var smallInt int8 = 100
	var bigInt int = int(smallInt)

	fmt.Printf("Value of bigInt is %d", bigInt)
}
