package user

import (
	"errors"
	U "final-project/entities/user"

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
	if err := repo.db.Create(&newUser).Error; err != nil {
		log.Warn(err)
		return U.Users{}, errors.New("gagal membuat user baru")
	}
	return newUser, nil
}

func (repo *UserRepository) Get(ID uint) (U.Users, error) {
	var user U.Users

	if err := repo.db.First(&user, ID).Error; err != nil {
		log.Warn(err)
		return U.Users{}, errors.New("data user tidak ditemukan")
	}
	return user, nil
}

func (repo *UserRepository) GetByID(ID uint) (U.Users, error) {
	var user U.Users

	if err := repo.db.First(&user, ID).Error; err != nil {
		log.Warn(err)
		return U.Users{}, errors.New("data user tidak ditemukan")
	}
	return user, nil
}

func (repo *UserRepository) GetAllUsers() ([]U.Users, error) {
	var users []U.Users
	repo.db.Find(&users)
	if len(users) < 1 {
		return nil, errors.New("tidak terdapat data user sama sekali")
	}
	return users, nil
}
