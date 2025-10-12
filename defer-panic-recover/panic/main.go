package main

import (
	"fmt"
	"os"
)

func loadConfig() string {
	config := os.Getenv("APP_CONFIG")
	if config == "" {
		panic("APP_CONFIG environment variable not set")
	}
	return config
}

func main() {
	defer fmt.Println("Shutting down...")
	fmt.Println("Loading configuration...")
	config := loadConfig()
	fmt.Println("Config loaded:", config)
}
