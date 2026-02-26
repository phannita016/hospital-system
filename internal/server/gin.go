package server

import (
	"hospital/config"
	"hospital/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewGinEngine(cfg *config.Config, db *gorm.DB) *gin.Engine {
	if cfg.Server.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery(), middleware.CORS(), middleware.Security())

	RegisterRoutes(router, db, cfg)

	return router
}
