package auth

import (
	"errors"
	U "final-project/entities/user"
	"final-project/repositories/hash"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) Login(email, password string) (U.Users, error) {
	var user U.Users

	repo.db.Model(&user).Where("email = ?", email).First(&user)

	isMatched := hash.CheckPasswordHash(user.Password, password)
	if !isMatched {
		return U.Users{}, errors.New("email dan password tidak cocok")
	}

	return user, nil
}
