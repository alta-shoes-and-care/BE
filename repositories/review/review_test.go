package review

import (
	"final-project/configs"
	O "final-project/entities/order"
	PM "final-project/entities/payment-method"
	R "final-project/entities/review"
	S "final-project/entities/service"
	U "final-project/entities/user"
	SeederOrder "final-project/repositories/mocks/order"
	SeederPaymentMethod "final-project/repositories/mocks/payment-method"
	SeederReview "final-project/repositories/mocks/review"
	SeederService "final-project/repositories/mocks/service"
	SeederUser "final-project/repositories/mocks/user"
	"final-project/repositories/order"
	paymentmethod "final-project/repositories/payment-method"
	"final-project/repositories/service"
	"final-project/repositories/user"
	"final-project/utils"
	"testing"

	"github.com/labstack/gommon/log"
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

func TestInsert(t *testing.T) {
	Migrator()
	userRepo := user.NewUserRepository(db)
	serviceRepo := service.NewServiceRepository(db)
	PMRepo := paymentmethod.NewPaymentMethodRepository(db)
	OrderRepo := order.NewOrderRepository(db)
	repo := NewReviewRepository(db)

	mockUser := SeederUser.UserSeeder()
	mockService := SeederService.ServiceSeeder()
	mockOrder := SeederOrder.OrderSeeder()
	mockPM := SeederPaymentMethod.PaymentMethodSeeder()
	mockReview := SeederReview.ReviewSeeder()

	t.Run("negative1", func(t *testing.T) {
		_, err := repo.Insert(mockReview)
		log.Info("ini gan", err)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		OrderRepo.Create(mockOrder)
		res, err := repo.Insert(mockReview)
		assert.Nil(t, err)
		assert.Equal(t, mockReview.Review, res.Review)
	})

	t.Run("negative2", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		OrderRepo.Create(mockOrder)
		_, err := repo.Insert(mockReview)
		assert.NotNil(t, err)
	})

}

func TestGet(t *testing.T) {
	Migrator()
	userRepo := user.NewUserRepository(db)
	serviceRepo := service.NewServiceRepository(db)
	PMRepo := paymentmethod.NewPaymentMethodRepository(db)
	OrderRepo := order.NewOrderRepository(db)
	repo := NewReviewRepository(db)

	mockUser := SeederUser.UserSeeder()
	mockService := SeederService.ServiceSeeder()
	mockOrder := SeederOrder.OrderSeeder()
	mockPM := SeederPaymentMethod.PaymentMethodSeeder()
	mockReview := SeederReview.ReviewSeeder()

	t.Run("negative", func(t *testing.T) {
		res, err := repo.Get()
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		OrderRepo.Create(mockOrder)
		repo.Insert(mockReview)
		res, err := repo.Get()
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}
