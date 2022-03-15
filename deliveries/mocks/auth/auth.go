package auth

import (
	"errors"
	U "final-project/entities/user"

	"gorm.io/gorm"
)

type MockAuthRepository struct{}

func (repo *MockAuthRepository) Login(email, password string) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Email: "ucup@ucup.com", Password: "ucup123"}, nil
}

type MockFalseAuthRepository struct{}

func (repo *MockFalseAuthRepository) Login(email, password string) (U.Users, error) {
	return U.Users{}, errors.New("false login")
}
