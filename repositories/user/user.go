package user

import (
	"errors"
	U "final-project/entities/user"
	"final-project/repositories/hash"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(newUser U.Users) (U.Users, error) {
	newUser.Password, _ = hash.HashPassword(newUser.Password)

	if err := repo.db.Create(&newUser).Error; err != nil {
		log.Warn(err)
		return U.Users{}, errors.New("gagal membuat user baru")
	}
	return newUser, nil
}

func (repo *UserRepository) Get(userID uint) (U.Users, error) {
	var user U.Users

	if err := repo.db.First(&user, userID).Error; err != nil {
		log.Warn(err)
		return U.Users{}, errors.New("data user tidak ditemukan")
	}
	return user, nil
}

func (repo *UserRepository) Update(userUpdate U.Users) (U.Users, error) {
	if userUpdate.Password != "" {
		userUpdate.Password, _ = hash.HashPassword(userUpdate.Password)
	}

	if rowsAffected := repo.db.Model(&userUpdate).Updates(userUpdate).Error; rowsAffected != nil {
		return U.Users{}, errors.New("tidak ada perubahan pada data user")
	}

	repo.db.First(&userUpdate)
	return userUpdate, nil
}

func (repo *UserRepository) Delete(userID uint) error {
	if rowsAffected := repo.db.Delete(&U.Users{}, userID).RowsAffected; rowsAffected == 0 {
		return errors.New("tidak ada user yang dihapus")
	}
	return nil
}
