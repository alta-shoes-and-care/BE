package auth

import (
	"errors"
	U "final-project/entities/user"

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

	if rowsAffected := repo.db.Model(&user).Where("email = ? AND password = ?", email, password).First(&user).RowsAffected; rowsAffected == 0 {
		return U.Users{}, errors.New("email dan password tidak cocok")
	}

	return user, nil
}
