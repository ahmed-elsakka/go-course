package main

import (
	"fmt"
	"time"
)

func fastOperation() {
	fmt.Println("Fast operation.")
}

func slowOperation() {
	time.Sleep(1 * time.Second)
	fmt.Println("Slow operation.")
}
func main() {
	start := time.Now()
	go slowOperation()
	go slowOperation()
	go slowOperation()
	go fastOperation()
	go fastOperation()
	fmt.Printf("Total run time: %v", time.Since(start))
}
