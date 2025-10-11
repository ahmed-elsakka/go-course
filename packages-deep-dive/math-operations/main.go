package main

import (
	"fmt"
	"math"
)

func main() {
	// Basic math operations using the math package

	// Constants
	fmt.Println("Pi:", math.Pi)
	fmt.Println("E :", math.E)

	// Power and Root
	fmt.Println("2^3:", math.Pow(2, 3))
	fmt.Println("Square Root of 16:", math.Sqrt(16))

	// Trigonometric functions (input in radians)
	fmt.Println("Sin(π/2):", math.Sin(math.Pi/2))
	fmt.Println("Cos(π):", math.Cos(math.Pi))
	fmt.Println("Tan(π/4):", math.Tan(math.Pi/4))

	// Logarithmic and exponential functions
	fmt.Println("Natural Log of 10:", math.Log(10))
	fmt.Println("Log base 10 of 100:", math.Log10(100))
	fmt.Println("e^2:", math.Exp(2))

	// Absolute value and rounding
	fmt.Println("Abs(-5.6):", math.Abs(-5.6))
	fmt.Println("Ceil(4.3):", math.Ceil(4.3))
	fmt.Println("Floor(4.7):", math.Floor(4.7))
	fmt.Println("Round(4.5):", math.Round(4.5))

	// Modulus and remainder
	fmt.Println("Mod(10, 3):", math.Mod(10, 3))

	// Max and Min
	fmt.Println("Max(10, 20):", math.Max(10, 20))
	fmt.Println("Min(10, 20):", math.Min(10, 20))
}
