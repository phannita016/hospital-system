package controller

import (
	"hospital/internal/modules/health/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
	healthService service.HealthService
}

func NewHealthController(healthService service.HealthService) *HealthController {
	return &HealthController{
		healthService: healthService,
	}
}

func (h *HealthController) CheckHealth(c *gin.Context) {
	res := h.healthService.CheckHealth()
	c.JSON(http.StatusOK, res)
}
