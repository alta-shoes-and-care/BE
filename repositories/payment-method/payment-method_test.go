package paymentmethod

import (
	"final-project/configs"
	O "final-project/entities/order"
	PM "final-project/entities/payment-method"
	R "final-project/entities/review"
	S "final-project/entities/service"
	U "final-project/entities/user"
	SeederPaymentMethod "final-project/seeders/payment-method"
	"final-project/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	config = configs.GetConfig(true)
	db     = utils.InitDB(config)
)

func Migrator() {
	db.Migrator().DropTable(&R.Reviews{})
	db.Migrator().DropTable(&O.Orders{})
	db.Migrator().DropTable(&S.Services{})
	db.Migrator().DropTable(&PM.PaymentMethods{})
	db.Migrator().DropTable(&U.Users{})

	db.AutoMigrate(&U.Users{})
	db.AutoMigrate(&PM.PaymentMethods{})
	db.AutoMigrate(&S.Services{})
	db.AutoMigrate(&O.Orders{})
	db.AutoMigrate(&R.Reviews{})
}

func TestCreate(t *testing.T) {
	Migrator()
	repo := NewPaymentMethodRepository(db)
	
	mockPM := SeederPaymentMethod.PaymentMethodSeeder()

	t.Run("positive", func(t *testing.T) {
		res, err := repo.Create(mockPM)
		assert.Nil(t, err)
		assert.Equal(t, mockPM.Name, res.Name)
	})

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Create(mockPM)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	Migrator()
	repo := NewPaymentMethodRepository(db)

	mockPM := SeederPaymentMethod.PaymentMethodSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Get()
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockPM)

		res, err := repo.Get()
		assert.Nil(t, err)
		assert.Equal(t, mockPM.Name, res[0].Name)
	})
}

func TestDelete(t *testing.T) {
	Migrator()
	repo := NewPaymentMethodRepository(db)

	mockPM := SeederPaymentMethod.PaymentMethodSeeder()

	t.Run("negative", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockPM)

		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
