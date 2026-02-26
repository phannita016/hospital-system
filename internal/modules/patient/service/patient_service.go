package service

import (
	"hospital/internal/modules/patient/entity"
	"hospital/internal/modules/patient/repository"
)

type PatientService interface {
	SearchPatients(hospitalID string, criteria map[string]interface{}) ([]entity.Patient, error)
}

type patientService struct {
	repo repository.PatientRepository
}

func NewPatientService(repo repository.PatientRepository) PatientService {
	return &patientService{repo: repo}
}

func (s *patientService) SearchPatients(hospitalID string, criteria map[string]interface{}) ([]entity.Patient, error) {
	return s.repo.Search(hospitalID, criteria)
}
