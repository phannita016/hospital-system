package main

import (
	"hospital/config"
	"hospital/internal/server"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server.Run(cfg)
}
