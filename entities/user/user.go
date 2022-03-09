package user

import (
	"final-project/entities/order"
	"final-project/entities/service"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string             `gorm:"type:varchar(255);not null"`
	Email    string             `gorm:"type:varchar(255);not null;unique"`
	Password string             `gorm:"type:varchar(255);not null"`
	IsAdmin  bool               `gorm:"type:boolean;default:false"`
	Services []service.Services `gorm:"foreignKey:UserID"`
	Orders   []order.Orders     `gorm:"foreignKey:UserID"`
}
