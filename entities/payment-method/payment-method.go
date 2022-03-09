package paymentmethod

import (
	"final-project/entities/order"

	"gorm.io/gorm"
)

type PaymentMethods struct {
	gorm.Model
	Name   string         `gorm:"type:varchar(255);not null"`
	Orders []order.Orders `gorm:"foreignKey:PaymentMethodID"`
}
