package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun and go is awesome"
	words := strings.Fields(text)
	countsMap := map[string]int{}

	for _, word := range words {
		countsMap[word]++
	}

	fmt.Println(countsMap)

}
