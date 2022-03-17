package midtranspay

import "github.com/midtrans/midtrans-go/coreapi"

type MidtransClient interface {
	CreateTransaction(orderID, bill uint) *coreapi.ChargeResponse
	CheckTransaction(orderID uint) (string, error)
}
