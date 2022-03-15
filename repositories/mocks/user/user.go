package user

import (
	U "final-project/entities/user"
)

func AdminSeeder() U.Users {
	mockAdmin := U.Users{
		Name:     "admin",
		Email:    "admin@mail.com",
		Password: "admin123",
		IsAdmin:  true,
	}
	return mockAdmin
}

func UserSeeder() U.Users {
	mockUser := U.Users{
		Name:     "ucup",
		Email:    "ucup@mail.com",
		Password: "ucup123",
		IsAdmin:  true,
	}
	return mockUser
}
