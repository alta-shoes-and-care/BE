package order

import (
	"errors"
	O "final-project/entities/order"
	orderRepo "final-project/repositories/order"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

const (
	layoutISO = "2006-01-02"
)

type MockTrueOrderRepository struct{}

func (repo *MockTrueOrderRepository) Create(newOrder O.Orders) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            false,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) Get() ([]orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")

	order1 := orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            false,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}
	order2 := orderRepo.FormatOrder{
		ID:                2,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            false,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-2",
	}

	return []orderRepo.FormatOrder{order1, order2}, nil
}

func (repo *MockTrueOrderRepository) GetByUserID(UserID uint) ([]orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")

	order1 := orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            false,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}
	order2 := orderRepo.FormatOrder{
		ID:                2,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            false,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-2",
	}

	return []orderRepo.FormatOrder{order1, order2}, nil
}

func (repo *MockTrueOrderRepository) GetByIDAdmin(ID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            false,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) GetByIDUser(ID, userID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            false,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) GetLastOrderID() (uint, error) {
	return 1, nil
}

func (repo *MockTrueOrderRepository) SetPaid(ID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            true,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) SetAccepted(ID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "accepted",
		IsPaid:            true,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) SetRejected(ID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "rejected",
		IsPaid:            false,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) SetOnProcess(ID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "on process",
		IsPaid:            true,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) SetDelivering(ID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "delivering",
		IsPaid:            true,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) SetCancelAdmin(ID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "cancel",
		IsPaid:            true,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) SetCancelUser(ID, userID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "cancel",
		IsPaid:            true,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) SetDone(ID, userID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "done",
		IsPaid:            true,
		HasRefunded:       false,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

func (repo *MockTrueOrderRepository) SetRefund(ID uint) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		UserName:          "Ucup",
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              date,
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "cancel",
		IsPaid:            true,
		HasRefunded:       true,
		CreatedAt:         time.Now(),
		Url:               "http://foo.com/bar-1",
	}, nil
}

type MockFalseOrderRepository struct{}

func (repo *MockFalseOrderRepository) Create(newOrder O.Orders) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to create")
}

func (repo *MockFalseOrderRepository) Get() ([]orderRepo.FormatOrder, error) {
	return nil, errors.New("fail to get")
}

func (repo *MockFalseOrderRepository) GetByUserID(UserID uint) ([]orderRepo.FormatOrder, error) {
	return nil, errors.New("fail to get by user id")
}

func (repo *MockFalseOrderRepository) GetByIDAdmin(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to get by id")
}

func (repo *MockFalseOrderRepository) GetByIDUser(ID, userID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to get by id")
}

func (repo *MockFalseOrderRepository) GetLastOrderID() (uint, error) {
	return 0, errors.New("no record")
}

func (repo *MockFalseOrderRepository) SetPaid(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set paid")
}

func (repo *MockFalseOrderRepository) SetAccepted(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set paid")
}

func (repo *MockFalseOrderRepository) SetRejected(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set rejected")
}

func (repo *MockFalseOrderRepository) SetOnProcess(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set on process")
}

func (repo *MockFalseOrderRepository) SetDelivering(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set delivering")
}

func (repo *MockFalseOrderRepository) SetCancelAdmin(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set cancel")
}

func (repo *MockFalseOrderRepository) SetCancelUser(ID, userID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set cancel")
}

func (repo *MockFalseOrderRepository) SetDone(ID, userID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set done")
}

func (repo *MockFalseOrderRepository) SetRefund(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set refund")
}

type MockTrueRedirectURLMidtrans struct{}

func (repo *MockTrueRedirectURLMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	return &coreapi.ChargeResponse{RedirectURL: "https://foo.com/bar-1"}, nil
}

func (repo *MockTrueRedirectURLMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "status pending", nil
}

type MockTrueActionsMidtrans struct{}

func (repo *MockTrueActionsMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	var a coreapi.Action
	a.URL = "https://foo.com/bar-1"

	return &coreapi.ChargeResponse{Actions: []coreapi.Action{a, a}}, nil
}

func (repo *MockTrueActionsMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "pending", nil
}

type MockFalseMidtrans struct{}

func (repo *MockFalseMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	return &coreapi.ChargeResponse{RedirectURL: ""}, errors.New("checking error")
}

func (repo *MockFalseMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "", errors.New("checking error")
}

type MockFalseRedirectURLMidtrans struct{}

func (repo *MockFalseRedirectURLMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	return &coreapi.ChargeResponse{RedirectURL: ""}, nil
}

func (repo *MockFalseRedirectURLMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "", errors.New("checking error")
}

type MockFalseActionsMidtrans struct{}

func (repo *MockFalseActionsMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	var a coreapi.Action
	a.URL = ""

	return &coreapi.ChargeResponse{Actions: []coreapi.Action{a, a}}, nil
}

func (repo *MockFalseActionsMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "", errors.New("checking error")
}

type MockSettlementMidtrans struct{}

func (repo *MockSettlementMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	return &coreapi.ChargeResponse{RedirectURL: "https://foo.com/bar-1"}, nil
}

func (repo *MockSettlementMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "status settlement", nil
}

type MockCancelMidtrans struct{}

func (repo *MockCancelMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	return &coreapi.ChargeResponse{RedirectURL: "https://foo.com/bar-1"}, nil
}

func (repo *MockCancelMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "status cancel", nil
}

type MockDenyMidtrans struct{}

func (repo *MockDenyMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	return &coreapi.ChargeResponse{RedirectURL: "https://foo.com/bar-1"}, nil
}

func (repo *MockDenyMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "status deny", nil
}

type MockAcceptMidtrans struct{}

func (repo *MockAcceptMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	return &coreapi.ChargeResponse{RedirectURL: "https://foo.com/bar-1"}, nil
}

func (repo *MockAcceptMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "status accept", nil
}

type MockChallengeMidtrans struct{}

func (repo *MockChallengeMidtrans) CreateTransaction(userID, orderID, bill uint, payment string) (*coreapi.ChargeResponse, error) {
	return &coreapi.ChargeResponse{RedirectURL: "https://foo.com/bar-1"}, nil
}

func (repo *MockChallengeMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "status challenge", nil
}
