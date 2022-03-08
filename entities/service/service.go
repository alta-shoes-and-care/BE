package service

import "gorm.io/gorm"

type Services struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text;not null"`
	Price       uint   `gorm:"not null"`
	Image       string `gorm:"type:text"`
	UserID      uint
}
