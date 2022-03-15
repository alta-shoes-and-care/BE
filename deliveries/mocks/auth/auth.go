package auth

import (
	"errors"
	U "final-project/entities/user"

	"gorm.io/gorm"
)

var (
	FalseJWT = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHBpcmVkIjoxNjQ3MjY5MzgzLCJpZCI6MSwiaXNBZG1pbiI6dHJ1ZX0.7pXr7rPxY_NGDmw09UycPXx--TzUkwO0_ZaveoBMQMQ"
)

type MockAuthRepository struct{}

func (repo *MockAuthRepository) Login(email, password string) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Email: "ucup@ucup.com", Password: "ucup123"}, nil
}

type MockFalseAuthRepository struct{}

func (repo *MockFalseAuthRepository) Login(email, password string) (U.Users, error) {
	return U.Users{}, errors.New("false login")
}
