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

func TestGet(t *testing.T) {
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
		res, err := repo.Get()
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.Get()
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestGetByUserID(t *testing.T) {
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
		res, err := repo.GetByUserID(1)
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.GetByUserID(1)
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestGetByID(t *testing.T) {
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
		_, err := repo.GetByID(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		_, err := repo.GetByID(1)
		assert.Nil(t, err)
	})
}

func TestInsertUrl(t *testing.T) {
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
		_, err := repo.InsertUrl(1, "")
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.InsertUrl(1, "https://foo.com/order/1")
		assert.Nil(t, err)
		assert.Equal(t, "https://foo.com/order/1", res.Url)
	})
}

func TestSetPaid(t *testing.T) {
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
		_, err := repo.SetPaid(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetPaid(1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.IsPaid)
	})
}

func TestSetAccepted(t *testing.T) {
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
		_, err := repo.SetAccepted(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetAccepted(1)
		assert.Nil(t, err)
		assert.Equal(t, "accepted", res.Status)
	})
}

func TestSetRejected(t *testing.T) {
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
		_, err := repo.SetRejected(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetRejected(1)
		assert.Nil(t, err)
		assert.Equal(t, "rejected", res.Status)
	})
}

func TestSetOnProcess(t *testing.T) {
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
		_, err := repo.SetOnProcess(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetOnProcess(1)
		assert.Nil(t, err)
		assert.Equal(t, "on process", res.Status)
	})
}

func TestSetDelivering(t *testing.T) {
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
		_, err := repo.SetDelivering(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetDelivering(1)
		assert.Nil(t, err)
		assert.Equal(t, "delivering", res.Status)
	})
}

func TestSetCancel(t *testing.T) {
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
		_, err := repo.SetCancel(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetCancel(1)
		assert.Nil(t, err)
		assert.Equal(t, "cancel", res.Status)
	})
}

func TestSetDone(t *testing.T) {
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
		_, err := repo.SetDone(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetDone(1)
		assert.Nil(t, err)
		assert.Equal(t, "done", res.Status)
	})
}