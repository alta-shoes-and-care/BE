package order

import O "final-project/entities/order"

type Order interface {
	Create(newOrder O.Orders) (FormatOrder, error)
	Get() ([]FormatOrder, error)
	GetByUserID(UserID uint) ([]FormatOrder, error)
	GetByID(ID, userID uint) (FormatOrder, error)
	GetLastOrderID() (uint, error)
	SetPaid(ID uint) (FormatOrder, error)
	SetAccepted(ID uint) (FormatOrder, error)
	SetRejected(ID uint) (FormatOrder, error)
	SetOnProcess(ID uint) (FormatOrder, error)
	SetDelivering(ID uint) (FormatOrder, error)
	SetCancel(ID uint) (FormatOrder, error)
	SetDone(ID uint) (FormatOrder, error)
}
