package main

import "fmt"

/*type Config struct {
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func main() {
	config := loadConfig("config.json")
	fmt.Println("Starting server on port:", config.Port)
	fmt.Println("Connected to database:", config.Database)
}

func loadConfig(filename string) Config {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file %s: %v", filename, err))
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(fmt.Sprintf("Invalid config format: %v", err))
	}

	return config
}*/

func main() {
	fmt.Println("Starting program...")
	panic("Unrecoverable issue happened")
	fmt.Println("This line wont execute")
}
