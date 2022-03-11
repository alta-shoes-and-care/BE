package order

import (
	"final-project/configs"
	O "final-project/entities/order"
	PM "final-project/entities/payment-method"
	R "final-project/entities/review"
	S "final-project/entities/service"
	U "final-project/entities/user"
	paymentmethod "final-project/repositories/payment-method"
	"final-project/repositories/service"
	"final-project/repositories/user"
	SeederOrder "final-project/seeders/order"
	SeederPaymentMethod "final-project/seeders/payment-method"
	SeederService "final-project/seeders/service"
	SeederUser "final-project/seeders/user"
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
	userRepo := user.NewUserRepository(db)
	serviceRepo := service.NewServiceRepository(db)
	PMRepo := paymentmethod.NewPaymentMethodRepository(db)
	repo := NewOrderRepository(db)

	mockUser := SeederUser.UserSeeder()
	mockService := SeederService.ServiceSeeder()
	mockOrder := SeederOrder.OrderSeeder()
	mockPM := SeederPaymentMethod.PaymentMethodSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Create(mockOrder)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		res, err := repo.Create(mockOrder)
		assert.Nil(t, err)
		assert.Equal(t, mockOrder.Phone, res.Phone)
	})
}
