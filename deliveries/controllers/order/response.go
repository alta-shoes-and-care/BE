package order

import (
	_OrderRepo "final-project/repositories/order"
)

var (
	layoutISOResponse = "2006-01-02 15:04:05"
)

type ResponseOrder struct {
	ID                uint   `json:"id"`
	UserID            uint   `json:"user_id"`
	UserName          string `json:"user_name"`
	ServiceID         uint   `json:"service_id"`
	ServiceTitle      string `json:"service_title"`
	Price             uint   `json:"price"`
	Qty               uint   `json:"qty"`
	PaymentMethodID   uint   `json:"payment_method_id"`
	PaymentMethodName string `json:"payment_method_name"`
	Date              string `json:"date"`
	Address           string `json:"address"`
	City              string `json:"city"`
	Phone             string `json:"phone"`
	Status            string `json:"status"`
	IsPaid            bool   `json:"is_paid"`
	CreatedAt         string `json:"created_at"`
	Url               string `json:"url"`
	Total             uint   `json:"total"`
}

func ToResponseOrder(order _OrderRepo.FormatOrder) ResponseOrder {
	return ResponseOrder{
		ID:                order.ID,
		UserID:            order.UserID,
		UserName:          order.UserName,
		ServiceID:         order.ServiceID,
		ServiceTitle:      order.ServiceTitle,
		Price:             order.Price,
		Qty:               order.Qty,
		PaymentMethodID:   order.PaymentMethodID,
		PaymentMethodName: order.PaymentMethodName,
		Date:              order.Date.Format(layoutISO),
		Address:           order.Address,
		City:              order.City,
		Phone:             order.Phone,
		Status:            order.Status,
		IsPaid:            order.IsPaid,
		CreatedAt:         order.CreatedAt.Format(layoutISOResponse),
		Url:               order.Url,

		Total: order.Price * order.Qty,
	}
}

func ToResponseOrderArr(orders []_OrderRepo.FormatOrder) []ResponseOrder {
	responses := make([]ResponseOrder, len(orders))

	for i := 0; i < len(orders); i++ {
		responses[i].ID = orders[i].ID
		responses[i].UserID = orders[i].UserID
		responses[i].UserName = orders[i].UserName
		responses[i].ServiceID = orders[i].ServiceID
		responses[i].ServiceTitle = orders[i].ServiceTitle
		responses[i].Price = orders[i].Price
		responses[i].Qty = orders[i].Qty
		responses[i].PaymentMethodID = orders[i].PaymentMethodID
		responses[i].PaymentMethodName = orders[i].PaymentMethodName
		responses[i].Date = orders[i].Date.Format(layoutISO)
		responses[i].Address = orders[i].Address
		responses[i].City = orders[i].City
		responses[i].Phone = orders[i].Phone
		responses[i].Status = orders[i].Status
		responses[i].IsPaid = orders[i].IsPaid
		responses[i].CreatedAt = orders[i].CreatedAt.Format(layoutISOResponse)
		responses[i].Url = orders[i].Url
		responses[i].Total = orders[i].Price * orders[i].Qty
	}

	return responses
}
