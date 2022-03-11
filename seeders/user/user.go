package user

import (
	"final-project/entities/user"
)

func AdminSeeder() user.Users {
	mockAdmin := user.Users{
		Name:     "admin",
		Email:    "admin@mail.com",
		Password: "admin123",
		IsAdmin:  true,
	}
	return mockAdmin
}

func UserSeeder() user.Users {
	mockUser := user.Users{
		Name:     "ucup",
		Email:    "ucup@mail.com",
		Password: "ucup123",
		IsAdmin:  true,
	}
	return mockUser
}
