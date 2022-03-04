package user

import (
	U "final-project/entities/user"

	"gorm.io/gorm"
)

type RequestCreateUser struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin"`
}

type ResponseCreateUser struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func (Req RequestCreateUser) ToEntityUser() U.Users {
	return U.Users{
		Name:     Req.Name,
		Email:    Req.Email,
		Password: Req.Password,
		IsAdmin:  Req.IsAdmin,
	}
}

func ToResponseCreateUser(User U.Users) ResponseCreateUser {
	return ResponseCreateUser{
		ID:      User.ID,
		Name:    User.Name,
		Email:   User.Email,
		IsAdmin: User.IsAdmin,
	}
}

type ResponseGetByID struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func ToResponseGetByID(User U.Users) ResponseGetByID {
	return ResponseGetByID{
		ID:      User.ID,
		Name:    User.Name,
		Email:   User.Email,
		IsAdmin: User.IsAdmin,
	}
}

type RequestUpdateUser struct {
	ID       uint
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin"`
}

type ResponseUpdateUser struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func (Req RequestUpdateUser) ToEntityUser(UserID uint) U.Users {
	return U.Users{
		Model:    gorm.Model{ID: UserID},
		Name:     Req.Name,
		Email:    Req.Email,
		Password: Req.Password,
		IsAdmin:  Req.IsAdmin,
	}
}

func ToResponseUpdate(User U.Users) ResponseUpdateUser {
	return ResponseUpdateUser{
		ID:      User.ID,
		Name:    User.Name,
		Email:   User.Email,
		IsAdmin: User.IsAdmin,
	}
}
