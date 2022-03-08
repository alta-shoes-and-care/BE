package paymentmethod

import (
	"errors"
	PM "final-project/entities/payment-method"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type PaymentMethodRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{
		db: db,
	}
}

func (repo *PaymentMethodRepository) Create(newPaymentMethod PM.PaymentMethods) (PM.PaymentMethods, error) {
	if err := repo.db.Create(&newPaymentMethod).Error; err != nil {
		log.Warn(err)
		return PM.PaymentMethods{}, errors.New("gagal membuat payment method baru")
	}
	return newPaymentMethod, nil
}

func (repo *PaymentMethodRepository) Get() ([]PM.PaymentMethods, error) {
	paymentmethods := []PM.PaymentMethods{}
	repo.db.Find(&paymentmethods)
	if len(paymentmethods) < 1 {
		return nil, errors.New("tidak terdapat payment method sama sekali")
	}
	return paymentmethods, nil
}

func (repo *PaymentMethodRepository) Delete(paymentMethodID uint) error {
	if rowsAffected := repo.db.Delete(&PM.PaymentMethods{}, paymentMethodID).RowsAffected; rowsAffected == 0 {
		return errors.New("tidak ada payment method yang dihapus")
	}
	return nil
}
