package paymentmethod

import (
	"errors"
	PM "final-project/entities/payment-method"
	U "final-project/entities/user"

	"gorm.io/gorm"
)

type MockAuthUserRepository struct{}

func (repo *MockAuthUserRepository) Login(email, password string) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Email: "ucup@ucup.com", Password: "ucup123", IsAdmin: false}, nil
}

type MockAuthAdminRepository struct{}

func (repo *MockAuthAdminRepository) Login(email, password string) (U.Users, error) {
	return U.Users{Model: gorm.Model{ID: 1}, Email: "ucup@ucup.com", Password: "ucup123", IsAdmin: true}, nil
}

type MockPaymentMethodRepository struct{}

func (repo *MockPaymentMethodRepository) Create(newPaymentMethod PM.PaymentMethods) (PM.PaymentMethods, error) {
	return PM.PaymentMethods{Model: gorm.Model{ID: 1}, Name: "klikbca"}, nil
}

func (repo *MockPaymentMethodRepository) Get() ([]PM.PaymentMethods, error) {
	payment1 := PM.PaymentMethods{
		Model:    gorm.Model{ID: 1},
		Name:     "klikbca",
	}

	payment2 := PM.PaymentMethods{
		Model:    gorm.Model{ID: 2},
		Name:     "gopay",
	}

	return []PM.PaymentMethods{payment1, payment2}, nil
}

func (repo *MockPaymentMethodRepository) Delete(paymentMethodID uint) error {
	return nil
}

type MockFalsePaymentMethodRepository struct{}

func (repo *MockFalsePaymentMethodRepository) Create(newPaymentMethod PM.PaymentMethods) (PM.PaymentMethods, error) {
	return PM.PaymentMethods{}, errors.New("false create payment method")
}

func (repo *MockFalsePaymentMethodRepository) Get() ([]PM.PaymentMethods, error) {
	return nil, errors.New("false get all payment methods")
}

func (repo *MockFalsePaymentMethodRepository) Delete(paymentMethodID uint) error {
	return errors.New("false delete payment method")
}