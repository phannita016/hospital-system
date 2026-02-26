package server

import (
	"hospital/internal/modules/health/controller"
	"hospital/internal/modules/health/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Initialize Health Module
	healthService := service.NewHealthService()
	healthController := controller.NewHealthController(healthService)

	// Health Routes
	router.GET("/health", healthController.CheckHealth)

	// Add other modules here...
}
