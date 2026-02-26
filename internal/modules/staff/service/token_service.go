package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(staffID uint, hospitalID string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type tokenService struct {
	secretKey []byte
}

func NewTokenService(secretKey string) TokenService {
	return &tokenService{secretKey: []byte(secretKey)}
}

type Claims struct {
	StaffID    uint   `json:"staff_id"`
	HospitalID string `json:"hospital_id"`
	jwt.RegisteredClaims
}

func (s *tokenService) GenerateToken(staffID uint, hospitalID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		StaffID:    staffID,
		HospitalID: hospitalID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *tokenService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
}
