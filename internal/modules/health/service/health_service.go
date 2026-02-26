package service

import "hospital/internal/modules/health/entity"

type HealthService interface {
	CheckHealth() *entity.Health
}

type healthService struct{}

func NewHealthService() HealthService {
	return &healthService{}
}

func (s *healthService) CheckHealth() *entity.Health {
	return &entity.Health{
		Status: "ok",
	}
}
