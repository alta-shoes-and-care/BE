package midtranspay

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func InitConnection() coreapi.Client {
	if err := godotenv.Load(".env"); err != nil {
		log.Info(err)
	}
	MIDTRANS_KEY := os.Getenv("MIDTRANS_KEY")

	client := coreapi.Client{}
	client.New(MIDTRANS_KEY, midtrans.Sandbox)
	return client
}

func CreateTransaction(client coreapi.Client, orderID, bill uint) *coreapi.ChargeResponse {
	req := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBCAKlikpay,
		BCAKlikPay: &coreapi.BCAKlikPayDetails{
			Desc: "Pembayaran BCA Klik Pay",
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("test-bayar-%d", orderID),
			GrossAmt: int64(bill),
		},
	}

	apiRes, err := client.ChargeTransaction(req)
	if err != nil {
		log.Warn("payment error:", err)
	}
	return apiRes
}

func CheckTransaction(client coreapi.Client, orderID uint) (string, error) {
	var result string

	transactionStatusResp, err := client.CheckTransaction(fmt.Sprintf("test-bayar-%d", orderID))
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
