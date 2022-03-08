package paymentmethod

import "gorm.io/gorm"

type PaymentMethods struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
}
