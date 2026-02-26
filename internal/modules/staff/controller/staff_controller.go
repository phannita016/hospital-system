package controller

import (
	"hospital/internal/modules/staff/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StaffController struct {
	staffService service.StaffService
	tokenService service.TokenService
}

func NewStaffController(staffService service.StaffService, tokenService service.TokenService) *StaffController {
	return &StaffController{
		staffService: staffService,
		tokenService: tokenService,
	}
}

type CreateStaffRequest struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	HospitalID string `json:"hospital" binding:"required"`
}

func (sc *StaffController) CreateStaff(c *gin.Context) {
	var req CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff, err := sc.staffService.CreateStaff(req.Username, req.Password, req.HospitalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, staff)
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (sc *StaffController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff, err := sc.staffService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := sc.tokenService.GenerateToken(staff.ID, staff.HospitalID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
