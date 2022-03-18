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

var (
	invoiceID string
)

func InitConnection() coreapi.Client {
	MIDTRANS_KEY := os.Getenv("MIDTRANS_KEY")

	client := coreapi.Client{}
	client.New(MIDTRANS_KEY, midtrans.Sandbox)
	return client
}

func (client *MidtransClientStruct) CreateTransaction(userID, orderID, bill uint) *coreapi.ChargeResponse {
	invoiceID = fmt.Sprintf("midtrans-test-%d%d", userID, orderID)

	req := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBCAKlikpay,
		BCAKlikPay: &coreapi.BCAKlikPayDetails{
			Desc: "Pembayaran BCA Klik Pay",
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  invoiceID,
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

func (client *MidtransClientStruct) CheckTransaction(userID, orderID uint) (string, error) {
	transactionStatusResp, err := client.CoreApiClient.CheckTransaction(invoiceID)
	if err != nil {
		log.Warn(err)
		return "", errors.New("gagal mengecek status pembayaran midtrans")
	} else {
		if transactionStatusResp != nil {
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					return "status challenge", nil
				} else if transactionStatusResp.FraudStatus == "accept" {
					return "status accept", nil
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				return "status settlement", nil
			} else if transactionStatusResp.TransactionStatus == "deny" {
				return "status deny", nil
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				return "status cancel", nil
			} else if transactionStatusResp.TransactionStatus == "pending" {
				return "status pending", nil
			}
		}
		return "", errors.New("respon status transaksi midtrans kosong")
	}
}
