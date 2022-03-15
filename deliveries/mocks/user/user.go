package user

import (
	"errors"
	U "final-project/entities/user"

	"gorm.io/gorm"
)

type MockAuthRepository struct{}

func (repo *MockAuthRepository) Login(email, password string) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Email: "ucup@ucup.com", Password: "ucup123"}, nil
}

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(newUser U.Users) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Name: "Yusuf Nur Wahid", Email: "ucup@ucup.com", Password: "ucup123", IsAdmin: true}, nil
}

func (repo *MockUserRepository) Get(ID uint) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Name: "Yusuf Nur Wahid", Email: "ucup@ucup.com", Password: "ucup123", IsAdmin: true}, nil
}

func (repo *MockUserRepository) GetByID(ID uint) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Name: "Yusuf Nur Wahid", Email: "ucup@ucup.com", Password: "ucup123", IsAdmin: true}, nil
}

func (repo *MockUserRepository) GetAllUsers() ([]U.Users, error) {
	user1 := U.Users{
		Model:    gorm.Model{ID: 1},
		Name:     "Yusuf",
		Email:    "ucup@ucup.com",
		Password: "ucup123",
		IsAdmin:  true,
	}

	user2 := U.Users{
		Model:    gorm.Model{ID: 2},
		Name:     "Frans",
		Email:    "fransp@ucup.com",
		Password: "frans123",
		IsAdmin:  true,
	}

	return []U.Users{user1, user2}, nil
}

func (repo *MockUserRepository) Update(userUpdate U.Users) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Name: "Yusuf Updated", Email: "ucup@ucup.com", Password: "ucup123", IsAdmin: true}, nil
}

func (repo *MockUserRepository) Delete(ID uint) error {
	return nil
}

type MockFalseUserRepository struct{}

func (repo *MockFalseUserRepository) Create(newUser U.Users) (U.Users, error) {
	return U.Users{}, errors.New("false create user")
}

func (repo *MockFalseUserRepository) Get(ID uint) (U.Users, error) {
	return U.Users{}, errors.New("false get user themselves")
}

func (repo *MockFalseUserRepository) GetByID(ID uint) (U.Users, error) {
	return U.Users{}, errors.New("false get user by id")
}

func (repo *MockFalseUserRepository) GetAllUsers() ([]U.Users, error) {
	return nil, errors.New("false get all users")
}

func (repo *MockFalseUserRepository) Update(userUpdate U.Users) (U.Users, error) {
	return U.Users{}, errors.New("false update user")
}

func (repo *MockFalseUserRepository) Delete(ID uint) error {
	return errors.New("false delete")
}
