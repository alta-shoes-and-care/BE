package order

import O "final-project/entities/order"

type Order interface {
	Create(newOrder O.Orders) (ResponseOrder, error)
	Get() ([]ResponseOrder, error)
	GetByUserID(UserID uint) ([]ResponseOrder, error)
	GetByID(ID uint) (ResponseOrder, error)
	InsertUrl(ID uint, url string) (ResponseOrder, error)
	SetPaid(ID uint) (ResponseOrder, error)
	SetAccepted(ID uint) (ResponseOrder, error)
	SetRejected(ID uint) (ResponseOrder, error)
	SetOnProcess(ID uint) (ResponseOrder, error)
	SetDelivering(ID uint) (ResponseOrder, error)
	SetCancel(ID uint) (ResponseOrder, error)
	SetDone(ID uint) (ResponseOrder, error)
}
