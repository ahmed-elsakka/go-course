package main

import (
	"fmt"
	"time"
)

func fastOperation(id int, channel chan string) {
	fmt.Printf("Fast operation %d started\n", id)
	channel <- fmt.Sprintf("Fast operation %d ended", id)
}

func slowOperation(id int, channel chan string) {
	fmt.Printf("Slow operation %d started\n", id)
	time.Sleep(1 * time.Second)
	channel <- fmt.Sprintf("Slow operation %d ended", id)
}
func main() {
	start := time.Now()
	channel := make(chan string)

	go slowOperation(1, channel)
	go slowOperation(2, channel)
	go slowOperation(3, channel)
	go fastOperation(4, channel)
	go fastOperation(5, channel)

	for i := 0; i < 4; i++ {
		fmt.Println(<-channel)
	}
	fmt.Printf("Total run time: %v", time.Since(start))
}
