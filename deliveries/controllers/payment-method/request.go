package paymentmethod

import (
	PM "final-project/entities/payment-method"
)

type RequestCreatePaymentMethod struct {
	Name     string `json:"name" form:"name"`
}

func (Req RequestCreatePaymentMethod) ToEntityPaymentMethod() PM.PaymentMethods {
	return PM.PaymentMethods{
		Name:     Req.Name,
	}
}