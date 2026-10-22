package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Items []string `json:"items"`
}

func loadConfig(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		return Config{}
	}
	defer file.Close()
	var cfg Config
	json.NewDecoder(file).Decode(&cfg)
	return cfg
}

func main() {
	cfg := loadConfig("config.json")
	fmt.Printf("Loaded %d items\n", len(cfg.Items))
}
