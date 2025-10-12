package main

import (
	"fmt"
	"time"
)

func fastOperation(id int, fastCh chan string) {
	fmt.Printf("Fast operation %d started\n", id)
	fastCh <- fmt.Sprintf("Fast operation %d ended", id)
}

func slowOperation(id int, slowCh chan string) {
	fmt.Printf("Slow operation %d started\n", id)
	time.Sleep(1 * time.Second)
	slowCh <- fmt.Sprintf("Slow operation %d ended", id)
}

func main() {
	start := time.Now()

	slowCh := make(chan string)
	fastCh := make(chan string)

	go slowOperation(1, slowCh)
	go slowOperation(2, slowCh)
	go slowOperation(3, slowCh)

	go fastOperation(4, fastCh)
	go fastOperation(5, fastCh)

	for i := 0; i < 5; i++ {
		select {
		case msg := <-slowCh:
			fmt.Println(msg)
		case msg := <-fastCh:
			fmt.Println(msg)
		}
	}

	fmt.Printf("Total run time: %v\n", time.Since(start))
}
