package paymentmethod

import (
	"final-project/entities/payment-method"
)

func PaymentMethodSeeder() paymentmethod.PaymentMethods {
	mockPaymentMethod := paymentmethod.PaymentMethods{
		Name:     "klikbca",
	}
	return mockPaymentMethod
}
