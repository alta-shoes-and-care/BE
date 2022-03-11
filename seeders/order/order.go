package order

import (
	"final-project/entities/order"
	"time"

	"gorm.io/datatypes"
)

func OrderSeeder() order.Orders {
	date, _ := time.Parse("2006-01-02", "2022-03-11")
	return order.Orders{
		Qty:             1,
		Date:            datatypes.Date(date),
		Address:         "Jl. Soedirman No.13",
		City:            "Surabaya",
		Url:             "https://blabla.com/tagihanbaru",
		Phone:           "081123456789",
		UserID:          1,
		ServiceID:       1,
		PaymentMethodID: 1,
	}
}
