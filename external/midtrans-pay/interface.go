package midtranspay

import "github.com/midtrans/midtrans-go/coreapi"

type MidtransClient interface {
	CreateTransaction(userID, orderID, bill uint) *coreapi.ChargeResponse
	CheckTransaction(userID, orderID uint) (string, error)
}
