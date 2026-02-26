package controller

import (
	"hospital/internal/modules/health/service"
	"hospital/internal/utils"

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
	utils.SuccessResponse(c, "", res)
}
