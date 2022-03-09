package service

import (
	"final-project/entities/order"

	"gorm.io/gorm"
)

type Services struct {
	gorm.Model
	Title       string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text;not null"`
	Price       uint           `gorm:"not null"`
	Image       string         `gorm:"type:text"`
	Orders      []order.Orders `gorm:"foreignKey:ServiceID"`
	UserID      uint
}
