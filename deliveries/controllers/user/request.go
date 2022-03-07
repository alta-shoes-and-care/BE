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

func (Req RequestCreateUser) ToEntityUser() U.Users {
	return U.Users{
		Name:     Req.Name,
		Email:    Req.Email,
		Password: Req.Password,
		IsAdmin:  Req.IsAdmin,
	}
}

type RequestUpdateUser struct {
	ID       uint
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	IsAdmin  bool   `json:"is_admin" form:"is_admin"`
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
