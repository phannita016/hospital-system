package server

import (
	"hospital/internal/modules/health/controller"
	"hospital/internal/modules/health/service"

	"hospital/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"hospital/internal/middleware"
	patientController "hospital/internal/modules/patient/controller"
	patientRepository "hospital/internal/modules/patient/repository"
	patientService "hospital/internal/modules/patient/service"
	staffController "hospital/internal/modules/staff/controller"
	staffRepository "hospital/internal/modules/staff/repository"
	staffService "hospital/internal/modules/staff/service"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, cfg *config.Config) {
	// Initialize Health Module
	healthService := service.NewHealthService()
	healthController := controller.NewHealthController(healthService)

	// Initialize Staff Module
	sRepo := staffRepository.NewStaffRepository(db)
	sService := staffService.NewStaffService(sRepo)
	tService := staffService.NewTokenService(cfg.Server.JWTSecret)
	sController := staffController.NewStaffController(sService, tService)

	// Initialize Patient Module
	pRepo := patientRepository.NewPatientRepository(db)
	pService := patientService.NewPatientService(pRepo)
	pController := patientController.NewPatientController(pService)

	// Health Routes
	router.GET("/health", healthController.CheckHealth)

	// Staff Routes
	router.POST("/staff/create", sController.CreateStaff)
	router.POST("/staff/login", sController.Login)

	// Patient Routes
	patientGroup := router.Group("/patient")
	patientGroup.Use(middleware.AuthMiddleware(tService))
	{
		patientGroup.GET("/search", pController.Search)
	}
}
