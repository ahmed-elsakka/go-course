package main

import (
	"fmt"
	"time"
)

func fast() {
	fmt.Println("Fast completed")
}

func slow() {
	time.Sleep(1 * time.Second)
	fmt.Println("Slow completed")
}

func main() {
	go slow()
	go fast()
	time.Sleep(3 * time.Second)
}
