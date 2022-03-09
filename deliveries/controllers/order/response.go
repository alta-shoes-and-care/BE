package order

import (
	_OrderRepo "final-project/repositories/order"
)

type ResponseOrder struct {
	_OrderRepo.FormatOrder
	Total uint `json:"total"`
}

func ToResponseOrder(order _OrderRepo.FormatOrder) ResponseOrder {
	return ResponseOrder{
		FormatOrder: _OrderRepo.FormatOrder{
			ID:                order.ID,
			UserID:            order.UserID,
			ServiceID:         order.ServiceID,
			ServiceTitle:      order.ServiceTitle,
			Price:             order.Price,
			Qty:               order.Qty,
			PaymentMethodID:   order.PaymentMethodID,
			PaymentMethodName: order.PaymentMethodName,
			Date:              order.Date,
			Address:           order.Address,
			City:              order.City,
			Phone:             order.Phone,
			Status:            order.Status,
			IsPaid:            order.IsPaid,
			Url:               order.Url,
		},
		Total: order.Price * order.Qty,
	}
}

func ToResponseOrderArr(orders []_OrderRepo.FormatOrder) []ResponseOrder {
	responses := make([]ResponseOrder, len(orders))

	for i := 0; i < len(orders); i++ {
		responses[i].ID = orders[i].ID
		responses[i].UserID = orders[i].UserID
		responses[i].ServiceID = orders[i].ServiceID
		responses[i].ServiceTitle = orders[i].ServiceTitle
		responses[i].Price = orders[i].Price
		responses[i].Qty = orders[i].Qty
		responses[i].PaymentMethodID = orders[i].PaymentMethodID
		responses[i].PaymentMethodName = orders[i].PaymentMethodName
		responses[i].Date = orders[i].Date
		responses[i].Address = orders[i].Address
		responses[i].City = orders[i].City
		responses[i].Phone = orders[i].Phone
		responses[i].Status = orders[i].Status
		responses[i].IsPaid = orders[i].IsPaid
		responses[i].Url = orders[i].Url
		responses[i].Total = orders[i].Price * orders[i].Qty
	}

	return responses
}
