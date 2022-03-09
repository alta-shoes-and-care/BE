package order

import (
	O "final-project/entities/order"
	"time"

	"gorm.io/datatypes"
)

type RequestCreateOrder struct {
	ServiceID       uint   `json:"service_id" form:"service_id"`
	Qty             uint   `json:"qty" form:"qty"`
	Total           uint   `json:"total" form:"total"`
	PaymentMethodID uint   `json:"payment_method_id" form:"payment_method_id"`
	Date            string `json:"date" form:"date"`
	Address         string `json:"address" form:"address"`
	City            string `json:"city" form:"city"`
	Phone           string `json:"phone" form:"phone"`
}

func (req RequestCreateOrder) ToEntityOrder(dateFormatted time.Time, userID uint) O.Orders {
	return O.Orders{
		Qty:             req.Qty,
		Date:            datatypes.Date(dateFormatted),
		Address:         req.Address,
		City:            req.City,
		Phone:           req.Phone,
		UserID:          userID,
		ServiceID:       req.ServiceID,
		PaymentMethodID: req.PaymentMethodID,
	}
}
