package order

import "time"

type FormatOrder struct {
	ID                uint
	UserID            uint
	UserName          string
	ServiceID         uint
	ServiceTitle      string
	Price             uint
	Qty               uint
	PaymentMethodID   uint
	PaymentMethodName string
	Date              time.Time
	Address           string
	City              string
	Phone             string
	Status            string
	IsPaid            bool
	HasRefunded       bool
	CreatedAt         time.Time
	Url               string
}

type LastOrderID struct {
	ID uint
}
