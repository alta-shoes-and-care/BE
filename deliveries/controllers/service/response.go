package service

import S "final-project/entities/service"

type ResponseCreate struct {
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Image       string `json:"image"`
}

func ToResponseCreate(service S.Services) ResponseCreate {
	return ResponseCreate{
		UserID:      service.UserID,
		Title:       service.Title,
		Description: service.Description,
		Price:       service.Price,
		Image:       service.Image,
	}
}

type ResponseGet struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Image       string `json:"image"`
}

func ToResponseGet(services []S.Services) []ResponseGet {
	responses := make([]ResponseGet, len(services))

	for i := 0; i < len(services); i++ {
		responses[i].ID = services[i].ID
		responses[i].UserID = services[i].UserID
		responses[i].Title = services[i].Title
		responses[i].Description = services[i].Description
		responses[i].Price = services[i].Price
		responses[i].Image = services[i].Image
	}

	return responses
}

func ToResponseGetDetails(service S.Services) ResponseGet {
	return ResponseGet{
		ID:          service.ID,
		UserID:      service.UserID,
		Title:       service.Title,
		Description: service.Description,
		Price:       service.Price,
		Image:       service.Image,
	}
}

type ResponseUpdate struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Image       string `json:"image"`
}

func ToResponseUpdate(service S.Services) ResponseUpdate {
	return ResponseUpdate{
		ID:          service.ID,
		UserID:      service.UserID,
		Title:       service.Title,
		Description: service.Description,
		Price:       service.Price,
		Image:       service.Image,
	}
}
