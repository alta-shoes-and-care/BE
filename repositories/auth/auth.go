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

func New(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (repo *AuthRepository) Login(email, password string) (U.Users, error) {
	var user U.Users

	repo.db.Model(&user).Where("email = ?", email).First(&user)

	isMatched := hash.CheckPasswordHash(user.Password, password)
	if !isMatched {
		return U.Users{}, errors.New("email atau password tidak valid")
	}

	return user, nil
}
