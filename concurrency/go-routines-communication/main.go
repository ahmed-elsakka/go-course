package main

import (
	"fmt"
	"time"
)

func fast(done chan int, slowDone chan int) {
	fmt.Println("Fast started")
	<-slowDone
	fmt.Println("Fast completed")
	done <- 1
}

func slow(done chan int) {
	fmt.Println("Slow started")
	time.Sleep(1 * time.Second)
	fmt.Println("Slow completed")
	done <- 1
}

func main() {
	slowDone := make(chan int)
	fastDone := make(chan int)

	go slow(slowDone)
	go fast(fastDone, slowDone)

	<-fastDone
}
