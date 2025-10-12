package main

import (
	"log"
	"os"
	"time"
)

func fastOperation(id int, fastCh chan string) {
	log.Printf("Fast operation %d started", id)
	fastCh <- "Fast operation ended"
}

func slowOperation(id int, slowCh chan string) {
	log.Printf("Slow operation %d started", id)
	time.Sleep(1 * time.Second)
	slowCh <- "Slow operation ended"
}

func main() {
	file, err := os.OpenFile("operations.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	start := time.Now()
	log.Println("Program started")

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
			log.Println(msg)
		case msg := <-fastCh:
			log.Println(msg)
		}
	}

	log.Printf("Total run time: %v", time.Since(start))
	log.Println("Program finished successfully")
}
