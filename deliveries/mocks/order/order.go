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

func (repo *MockTrueOrderRepository) GetByID(ID uint) (orderRepo.FormatOrder, error) {
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

func (repo *MockTrueOrderRepository) SetCancel(ID uint) (orderRepo.FormatOrder, error) {
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

func (repo *MockTrueOrderRepository) SetDone(ID uint) (orderRepo.FormatOrder, error) {
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

func (repo *MockFalseOrderRepository) GetByID(ID uint) (orderRepo.FormatOrder, error) {
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

func (repo *MockFalseOrderRepository) SetCancel(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set cancel")
}

func (repo *MockFalseOrderRepository) SetDone(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set done")
}

func (repo *MockFalseOrderRepository) SetRefund(ID uint) (orderRepo.FormatOrder, error) {
	return orderRepo.FormatOrder{}, errors.New("fail to set refund")
}

type MockTrueMidtrans struct{}

func (repo *MockTrueMidtrans) CreateTransaction(userID, orderID, bill uint) *coreapi.ChargeResponse {
	return &coreapi.ChargeResponse{
		RedirectURL: "https://foo.com/bar-1",
	}
}

func (repo *MockTrueMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "status pending", nil
}

type MockFalseMidtrans struct{}

func (repo *MockFalseMidtrans) CreateTransaction(userID, orderID, bill uint) *coreapi.ChargeResponse {
	return &coreapi.ChargeResponse{
		RedirectURL: "",
	}
}

func (repo *MockFalseMidtrans) CheckTransaction(userID, orderID uint) (string, error) {
	return "", errors.New("checking error")
}
