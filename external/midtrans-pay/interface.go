package midtranspay

import "github.com/midtrans/midtrans-go/coreapi"

type MidtransClient interface {
	CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error)
	CheckTransaction(userID, orderID uint) (string, error)
}
