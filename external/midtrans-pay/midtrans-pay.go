package midtranspay

import (
	"errors"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type MidtransClientStruct struct {
	CoreApiClient coreapi.Client
}

func NewMidtransClient(client coreapi.Client) *MidtransClientStruct {
	return &MidtransClientStruct{
		CoreApiClient: client,
	}
}

func InitConnection() coreapi.Client {
	MIDTRANS_KEY := os.Getenv("MIDTRANS_KEY")

	client := coreapi.Client{}
	client.New(MIDTRANS_KEY, midtrans.Sandbox)
	return client
}

func (client *MidtransClientStruct) CreateTransaction(orderID, bill uint) *coreapi.ChargeResponse {
	req := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBCAKlikpay,
		BCAKlikPay: &coreapi.BCAKlikPayDetails{
			Desc: "Pembayaran BCA Klik Pay",
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("test-bayar2-%d", orderID),
			GrossAmt: int64(bill),
		},
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: 5,
			Unit:           "minute",
		},
	}

	apiRes, err := client.CoreApiClient.ChargeTransaction(req)
	if err != nil {
		log.Warn("payment error:", err)
	}
	return apiRes
}

func (client *MidtransClientStruct) CheckTransaction(orderID uint) (string, error) {
	var result string

	transactionStatusResp, err := client.CoreApiClient.CheckTransaction(fmt.Sprintf("test-bayar-%d", orderID))
	if err != nil {
		log.Warn(err)
		return "", errors.New("gagal mengecek status transaksi")
	} else {
		if transactionStatusResp != nil {
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					result = "status challenge"
				} else if transactionStatusResp.FraudStatus == "accept" {
					result = "status accept"
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				result = "status settlement"
			} else if transactionStatusResp.TransactionStatus == "deny" {
				result = "status deny"
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				result = "status cancel"
			} else if transactionStatusResp.TransactionStatus == "pending" {
				result = "status pending"
			}
		}
	}
	return result, nil
}
