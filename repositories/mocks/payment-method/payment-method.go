package paymentmethod

import (
	PM "final-project/entities/payment-method"
)

func PaymentMethodSeeder() PM.PaymentMethods {
	mockPaymentMethod := PM.PaymentMethods{
		Name: "klikbca",
	}
	return mockPaymentMethod
}
