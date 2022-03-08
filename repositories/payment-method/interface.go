package paymentmethod

import PM "final-project/entities/payment-method"

type PaymentMethod interface {
	Create(newPaymentMethod PM.PaymentMethods) (PM.PaymentMethods, error)
	Get() ([]PM.PaymentMethods, error)
	Delete(paymentMethodID uint) error
}
