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

func (client *MidtransClientStruct) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	invoiceID = fmt.Sprintf("midtrans-testxxx-%d%d", userID, orderID)
	var req *coreapi.ChargeReq

	trxDetails := midtrans.TransactionDetails{
		OrderID:  invoiceID,
		GrossAmt: int64(bill),
	}

	expiry := coreapi.CustomExpiry{
		ExpiryDuration: 5,
		Unit:           "minute",
	}

	switch payment {
	case "bca klikpay":
		req = &coreapi.ChargeReq{
			PaymentType: coreapi.PaymentTypeBCAKlikpay,
			BCAKlikPay: &coreapi.BCAKlikPayDetails{
				Desc: "Pembayaran BCA Klik Pay",
			},
			TransactionDetails: trxDetails,
			CustomExpiry:       &expiry,
		}
	case "klikbca":
		req = &coreapi.ChargeReq{
			PaymentType: coreapi.PaymentTypeKlikBca,
			BCAKlikBCA: &coreapi.BcaKlikBCADetails{
				Desc:   "Pembayaran Klik BCA",
				UserID: fmt.Sprintf("%d", userID),
			},
			TransactionDetails: trxDetails,
			CustomExpiry:       &expiry,
		}
	case "gopay":
		req = &coreapi.ChargeReq{
			PaymentType:        coreapi.PaymentTypeGopay,
			TransactionDetails: trxDetails,
			CustomExpiry:       &expiry,
		}
	case "cimb clicks":
		req = &coreapi.ChargeReq{
			PaymentType: coreapi.PaymentTypeCimbClicks,
			CIMBClicks: &coreapi.CIMBClicksDetails{
				Desc: "Pembayaran CIMB Clicks",
			},
			TransactionDetails: trxDetails,
			CustomExpiry:       &expiry,
		}
	case "bri epay":
		req = &coreapi.ChargeReq{
			PaymentType:        coreapi.PaymentTypeBRIEpay,
			TransactionDetails: trxDetails,
			CustomExpiry:       &expiry,
		}
	}

	apiRes, err := client.CoreApiClient.ChargeTransaction(req)
	if err != nil {
		log.Warn("payment error:", err)
		return nil, errors.New("gagal melakukan charge transaction midtrans")
	}
	return apiRes, nil
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
