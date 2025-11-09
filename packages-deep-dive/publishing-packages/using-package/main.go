package main

import (
	"fmt"

	"github.com/ahmed-elsakka/mathutils"
)

func main() {
	fmt.Println(mathutils.Add(1, 2))
	fmt.Println(mathutils.Multiply(1, 2))
	fmt.Println(mathutils.Average([]float64{2, 4, 6}))
}
