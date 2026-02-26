package server

import (
	"hospital/config"
	"hospital/internal/driver"
	"log"
)

func Run(cfg *config.Config) {
	db := cfg.Database

	pool, _, err := driver.NewDB(db)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("Database connection successful. Application is running.")

	engine := NewGinEngine(cfg)

	port := cfg.Server.Port
	log.Printf("Starting server on %s", port)
	if err := engine.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
