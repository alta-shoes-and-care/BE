package user

import (
	"final-project/entities/user"
	"final-project/repositories/hash"
)

func AdminSeeder() user.Users {
	password, _ := hash.HashPassword("admin123")

	mockAdmin := user.Users{
		Name:     "admin",
		Email:    "admin@mail.com",
		Password: password,
		IsAdmin:  true,
	}
	return mockAdmin
}

func UserSeeder() user.Users {
	password, _ := hash.HashPassword("ucup123")

	mockUser := user.Users{
		Name:     "ucup",
		Email:    "ucup@mail.com",
		Password: password,
		IsAdmin:  true,
	}
	return mockUser
}
