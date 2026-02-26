package repository

import (
	"hospital/internal/modules/patient/entity"

	"gorm.io/gorm"
)

type PatientRepository interface {
	Search(hospitalID string, criteria map[string]interface{}) ([]entity.Patient, error)
}

type patientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) Search(hospitalID string, criteria map[string]interface{}) ([]entity.Patient, error) {
	var patients []entity.Patient
	query := r.db.Where("hospital_id = ?", hospitalID)

	if val, ok := criteria["national_id"]; ok && val != "" {
		query = query.Where("national_id = ?", val)
	}
	if val, ok := criteria["passport_id"]; ok && val != "" {
		query = query.Where("passport_id = ?", val)
	}
	if val, ok := criteria["first_name"]; ok && val != "" {
		query = query.Where("first_name_th LIKE ? OR first_name_en LIKE ?", "%"+val.(string)+"%", "%"+val.(string)+"%")
	}
	if val, ok := criteria["last_name"]; ok && val != "" {
		query = query.Where("last_name_th LIKE ? OR last_name_en LIKE ?", "%"+val.(string)+"%", "%"+val.(string)+"%")
	}
	// Add more search criteria as needed...

	if err := query.Find(&patients).Error; err != nil {
		return nil, err
	}
	return patients, nil
}
