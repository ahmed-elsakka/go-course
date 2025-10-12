package main

import (
	"fmt"

	"testing-example/mathutils"
)

func main() {
	fmt.Printf("1 + 2 = %d", mathutils.Add(1, 2))
	fmt.Printf("1 * 2 = %d", mathutils.Multiply(1, 2))
}
