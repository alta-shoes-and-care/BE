package user

import U "final-project/entities/user"

type ResponseCreateUser struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func ToResponseCreateUser(User U.Users) ResponseCreateUser {
	return ResponseCreateUser{
		ID:      User.ID,
		Name:    User.Name,
		Email:   User.Email,
		IsAdmin: User.IsAdmin,
	}
}

type ResponseGetUser struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func ToResponseGetUser(User U.Users) ResponseGetUser {
	return ResponseGetUser{
		ID:      User.ID,
		Name:    User.Name,
		Email:   User.Email,
		IsAdmin: User.IsAdmin,
	}
}

type ResponseUpdateUser struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
}

func ToResponseUpdateUser(User U.Users) ResponseUpdateUser {
	return ResponseUpdateUser{
		ID:      User.ID,
		Name:    User.Name,
		Email:   User.Email,
		IsAdmin: User.IsAdmin,
	}
}
