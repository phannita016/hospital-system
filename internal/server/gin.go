package server

import (
	"hospital/config"
	"hospital/internal/middleware"

	"github.com/gin-gonic/gin"
)

func NewGinEngine(cfg *config.Config) *gin.Engine {
	if cfg.Server.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery(), middleware.CORS(), middleware.Security())

	RegisterRoutes(router)

	return router
}
