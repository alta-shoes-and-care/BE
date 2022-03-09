package order

import "time"

type ResponseOrder struct {
	ID                uint      `json:"id"`
	UserID            uint      `json:"user_id"`
	ServiceID         uint      `json:"service_id"`
	ServiceTitle      string    `json:"service_title"`
	Price             uint      `json:"price"`
	Qty               uint      `json:"qty"`
	PaymentMethodID   uint      `json:"payment_method_id"`
	PaymentMethodName string    `json:"payment_method_name"`
	Date              time.Time `json:"date"`
	Address           string    `json:"address"`
	City              string    `json:"city"`
	Phone             string    `json:"phone"`
	Status            string    `json:"status"`
	IsPaid            bool      `json:"is_paid"`
	Url               string    `json:"url"`
}