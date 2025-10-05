package main

import "fmt"

type DivisionError struct {
	A, B float64
	Msg  string
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("Math error: %s (a=%.2f, b=%.2f)", e.Msg, e.A, e.B)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &DivisionError{A: a, B: b, Msg: "cannot divide by zero"}
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Result: ", result)
}
