package user

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null"`
	IsAdmin  bool   `gorm:"type:boolean;default:false"`
}
