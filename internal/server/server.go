package server

import (
	"hospital/config"
	"hospital/internal/driver"
	"log"
)

func Run(cfg *config.Config) {
	db, cleanup, err := driver.NewDB(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer cleanup()
	_ = db // Suppress unused variable error until repositories are implemented
	log.Println("Database connection successful. Application is running.")

	engine := NewGinEngine(cfg)

	port := cfg.Server.Port
	log.Printf("Starting server on %s", port)
	if err := engine.Run(port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
