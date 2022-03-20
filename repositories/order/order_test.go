package order

import (
	"final-project/configs"
	O "final-project/entities/order"
	PM "final-project/entities/payment-method"
	R "final-project/entities/review"
	S "final-project/entities/service"
	U "final-project/entities/user"
	SeederOrder "final-project/repositories/mocks/order"
	SeederPaymentMethod "final-project/repositories/mocks/payment-method"
	SeederService "final-project/repositories/mocks/service"
	SeederUser "final-project/repositories/mocks/user"
	paymentmethod "final-project/repositories/payment-method"
	"final-project/repositories/service"
	"final-project/repositories/user"
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

func TestGetByIDAdmin(t *testing.T) {
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
		_, err := repo.GetByIDAdmin(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		_, err := repo.GetByIDAdmin(1)
		assert.Nil(t, err)
	})
}

func TestGetByIDUser(t *testing.T) {
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
		_, err := repo.GetByIDUser(1, 1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		_, err := repo.GetByIDUser(1, 1)
		assert.Nil(t, err)
	})
}

func TestGetLastOrderID(t *testing.T) {
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
		_, err := repo.GetLastOrderID()
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.GetLastOrderID()
		assert.Nil(t, err)
		assert.NotZero(t, res)
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

func TestSetCancelAdmin(t *testing.T) {
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
		_, err := repo.SetCancelAdmin(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetCancelAdmin(1)
		assert.Nil(t, err)
		assert.Equal(t, "cancel", res.Status)
	})
}

func TestSetCancelUser(t *testing.T) {
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
		_, err := repo.SetCancelUser(1, 1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetCancelUser(1, 1)
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
		_, err := repo.SetDone(1, 1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)
		repo.Create(mockOrder)
		res, err := repo.SetDone(1, 1)
		assert.Nil(t, err)
		assert.Equal(t, "done", res.Status)
	})
}

func TestSetRefund(t *testing.T) {
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
		_, err := repo.SetRefund(1)
		assert.NotNil(t, err)
	})

	t.Run("positive 1", func(t *testing.T) {
		userRepo.Create(mockUser)
		serviceRepo.Create(mockService)
		PMRepo.Create(mockPM)

		mockOrder.Status = "cancel"
		mockOrder.IsPaid = true
		repo.Create(mockOrder)

		res, err := repo.SetRefund(1)
		assert.Nil(t, err)
		assert.Equal(t, true, res.HasRefunded)
	})

	t.Run("positive 2", func(t *testing.T) {
		mockOrder.Status = "rejected"
		mockOrder.IsPaid = true
		repo.Create(mockOrder)

		res, err := repo.SetRefund(2)
		assert.Nil(t, err)
		assert.Equal(t, true, res.HasRefunded)
	})
}
