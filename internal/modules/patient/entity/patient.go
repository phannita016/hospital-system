package entity

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	NationalID   string         `gorm:"index" json:"national_id"`
	PassportID   string         `gorm:"index" json:"passport_id"`
	FirstNameTH  string         `json:"first_name_th"`
	MiddleNameTH string         `json:"middle_name_th"`
	LastNameTH   string         `json:"last_name_th"`
	FirstNameEN  string         `json:"first_name_en"`
	MiddleNameEN string         `json:"middle_name_en"`
	LastNameEN   string         `json:"last_name_en"`
	DOB          time.Time      `json:"date_of_birth"`
	PhoneNumber  string         `json:"phone_number"`
	Email        string         `json:"email"`
	Gender       string         `json:"gender"`
	HospitalHN   string         `json:"patient_hn"`
	HospitalID   string         `gorm:"index" json:"hospital_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
