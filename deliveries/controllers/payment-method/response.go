package paymentmethod

import PM "final-project/entities/payment-method"

type ResponseCreatePaymentMethod struct {
	Name string `json:"name"`
}

func ToResponseCreatePaymentMethod(PaymentMethod PM.PaymentMethods) ResponseCreatePaymentMethod {
	return ResponseCreatePaymentMethod{
		Name: PaymentMethod.Name,
	}
}

type ResponseGetPaymentMethod struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func ToResponseGetPaymentMethod(PaymentMethod []PM.PaymentMethods) []ResponseGetPaymentMethod {
	responses := make([]ResponseGetPaymentMethod, len(PaymentMethod))
	for i := 0; i < len(PaymentMethod); i++ {
		responses[i].ID = PaymentMethod[i].ID
		responses[i].Name = PaymentMethod[i].Name
	}
	return responses
}
