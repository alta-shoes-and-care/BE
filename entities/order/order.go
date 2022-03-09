package order

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Qty             uint
	Date            time.Time
	Address         string `gorm:"type:text;not null"`
	City            string `gorm:"type:varchar(255);not null"`
	Status          string `gorm:"type:enum('pending', 'cancel', 'accepted', 'rejected', 'on process', 'delivering', 'done');default:'pending'"`
	IsPaid          bool   `gorm:"type:boolean;default:false"`
	Url             string
	Phone           string
	UserID          uint
	ServiceID       uint
	PaymentMethodID uint
}
