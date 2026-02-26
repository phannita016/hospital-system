package service

import (
	"errors"
	"hospital/internal/modules/staff/entity"
	"hospital/internal/modules/staff/repository"

	"golang.org/x/crypto/bcrypt"
)

type StaffService interface {
	CreateStaff(username, password, hospitalID string) (*entity.Staff, error)
	Login(username, password string) (*entity.Staff, error)
}

type staffService struct {
	repo repository.StaffRepository
}

func NewStaffService(repo repository.StaffRepository) StaffService {
	return &staffService{repo: repo}
}

func (s *staffService) CreateStaff(username, password, hospitalID string) (*entity.Staff, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	staff := &entity.Staff{
		Username:   username,
		Password:   string(hashedPassword),
		HospitalID: hospitalID,
	}

	if err := s.repo.Create(staff); err != nil {
		return nil, err
	}

	return staff, nil
}

func (s *staffService) Login(username, password string) (*entity.Staff, error) {
	staff, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid username or password")
	}

	return staff, nil
}
