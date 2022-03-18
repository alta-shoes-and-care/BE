package order

import O "final-project/entities/order"

type Order interface {
	Create(newOrder O.Orders) (FormatOrder, error)
	Get() ([]FormatOrder, error)
	GetByUserID(UserID uint) ([]FormatOrder, error)
	GetByIDAdmin(ID uint) (FormatOrder, error)
	GetByIDUser(ID, userID uint) (FormatOrder, error)
	GetLastOrderID() (uint, error)
	SetPaid(ID uint) (FormatOrder, error)
	SetAccepted(ID uint) (FormatOrder, error)
	SetRejected(ID uint) (FormatOrder, error)
	SetOnProcess(ID uint) (FormatOrder, error)
	SetDelivering(ID uint) (FormatOrder, error)
	SetCancelAdmin(ID uint) (FormatOrder, error)
	SetCancelUser(ID, userID uint) (FormatOrder, error)
	SetDone(ID, userID uint) (FormatOrder, error)
	SetRefund(ID uint) (FormatOrder, error)
}
