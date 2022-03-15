package service

import S "final-project/entities/service"

func ServiceSeeder() S.Services {
	mockService := S.Services{
		Title:       "service 1",
		Description: "layanan 1",
		Price:       10000,
		Image:       "https://blabla.com/",
		UserID:      1,
	}
	return mockService
}
