package entity

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Username   string         `gorm:"unique;not null" json:"username"`
	Password   string         `gorm:"not null" json:"-"`
	HospitalID string         `gorm:"not null" json:"hospital_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}
