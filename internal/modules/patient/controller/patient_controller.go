package controller

import (
	"hospital/internal/modules/patient/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientController struct {
	patientService service.PatientService
}

func NewPatientController(patientService service.PatientService) *PatientController {
	return &PatientController{
		patientService: patientService,
	}
}

func (pc *PatientController) Search(c *gin.Context) {
	hospitalID, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "hospital context not found"})
		return
	}

	criteria := make(map[string]interface{})
	if nationalID := c.Query("national_id"); nationalID != "" {
		criteria["national_id"] = nationalID
	}
	if passportID := c.Query("passport_id"); passportID != "" {
		criteria["passport_id"] = passportID
	}
	if firstName := c.Query("first_name"); firstName != "" {
		criteria["first_name"] = firstName
	}
	if lastName := c.Query("last_name"); lastName != "" {
		criteria["last_name"] = lastName
	}

	patients, err := pc.patientService.SearchPatients(hospitalID.(string), criteria)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}
