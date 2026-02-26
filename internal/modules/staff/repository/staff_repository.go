package repository

import (
	"hospital/internal/modules/staff/entity"

	"gorm.io/gorm"
)

type StaffRepository interface {
	Create(staff *entity.Staff) error
	FindByUsername(username string) (*entity.Staff, error)
	FindByID(id uint) (*entity.Staff, error)
}

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &staffRepository{db: db}
}

func (r *staffRepository) Create(staff *entity.Staff) error {
	return r.db.Create(staff).Error
}

func (r *staffRepository) FindByUsername(username string) (*entity.Staff, error) {
	var staff entity.Staff
	if err := r.db.Where("username = ?", username).First(&staff).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r *staffRepository) FindByID(id uint) (*entity.Staff, error) {
	var staff entity.Staff
	if err := r.db.First(&staff, id).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}
