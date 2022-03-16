package order

import (
	O "final-project/entities/order"
	orderRepo "final-project/repositories/order"
	"time"

	"gorm.io/datatypes"
)

const (
	layoutISO = "2006-01-02"
)

type MockTrueOrderRepository struct{}

func (repo *MockTrueOrderRepository) Create(newOrder O.Orders) (orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")
	return orderRepo.FormatOrder{ID: 1, UserID: 1, ServiceID: 1, ServiceTitle: "Regular Cleaning", Price: 15000, Qty: 1, PaymentMethodID: 1, PaymentMethodName: "BCA Klikpay", Date: datatypes.Date(date), Address: "Jl. Soedirman", City: "Jakarta", Phone: "080000000000", Status: "pending", IsPaid: false, Url: "http://foo.com/bar-1"}, nil
}

func (repo *MockTrueOrderRepository) Get() ([]orderRepo.FormatOrder, error) {
	date, _ := time.Parse(layoutISO, "2022-03-16")

	order1 := orderRepo.FormatOrder{
		ID:                1,
		UserID:            1,
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              datatypes.Date(date),
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            false,
		Url:               "http://foo.com/bar-1",
	}
	order2 := orderRepo.FormatOrder{
		ID:                2,
		UserID:            1,
		ServiceID:         1,
		ServiceTitle:      "Regular Cleaning",
		Price:             15000,
		Qty:               1,
		PaymentMethodID:   1,
		PaymentMethodName: "BCA Klikpay",
		Date:              datatypes.Date(date),
		Address:           "Jl. Soedirman",
		City:              "Jakarta",
		Phone:             "080000000000",
		Status:            "pending",
		IsPaid:            false,
		Url:               "http://foo.com/bar-2",
	}

	return []orderRepo.FormatOrder{order1, order2}, nil
}

func (repo *MockTrueOrderRepository) GetByUserID(UserID uint) ([]orderRepo.FormatOrder, error) {}

func (repo *MockTrueOrderRepository) GetByID(ID uint) (orderRepo.FormatOrder, error) {}

func (repo *MockTrueOrderRepository) GetLastOrderID() (uint, error) {}

func (repo *MockTrueOrderRepository) SetPaid(ID uint) (orderRepo.FormatOrder, error) {}

func (repo *MockTrueOrderRepository) SetAccepted(ID uint) (orderRepo.FormatOrder, error) {}

func (repo *MockTrueOrderRepository) SetRejected(ID uint) (orderRepo.FormatOrder, error) {}

func (repo *MockTrueOrderRepository) SetOnProcess(ID uint) (orderRepo.FormatOrder, error) {}

func (repo *MockTrueOrderRepository) SetDelivering(ID uint) (orderRepo.FormatOrder, error) {}

func (repo *MockTrueOrderRepository) SetCancel(ID uint) (orderRepo.FormatOrder, error) {}

func (repo *MockTrueOrderRepository) SetDone(ID uint) (orderRepo.FormatOrder, error) {}
