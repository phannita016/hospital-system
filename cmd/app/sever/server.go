package sever

import (
	"hospital/internal/configs"
	"hospital/internal/driver"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	log.Println("Attempting to connect to the database...")
	if err := configs.Setup(); err != nil {
		log.Fatalf("config setup failed: %v", err)
	}

	dbConfig := configs.Get().Database
	pool, _, err := driver.NewDB(dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("Database connection successful. Application is running.")

	router := gin.Default()

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
