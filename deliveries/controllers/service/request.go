package service

import (
	S "final-project/entities/service"

	"gorm.io/gorm"
)

type RequestCreate struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Price       uint   `json:"price" form:"price"`
}

func (Req RequestCreate) ToEntityService(image string, userID uint) S.Services {
	return S.Services{
		Title:       Req.Title,
		Description: Req.Description,
		Price:       Req.Price,
		Image:       image,
		UserID:      userID,
	}
}

type RequestUpdate struct {
	ID          uint   `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Price       uint   `json:"price" form:"price"`
}

func (Req RequestUpdate) ToEntityService(image string) S.Services {
	return S.Services{
		Model:       gorm.Model{ID: Req.ID},
		Title:       Req.Title,
		Description: Req.Description,
		Price:       Req.Price,
		Image:       image,
		UserID:      0,
	}
}
