package service

import (
	"final-project/configs"
	O "final-project/entities/order"
	PM "final-project/entities/payment-method"
	R "final-project/entities/review"
	S "final-project/entities/service"
	U "final-project/entities/user"
	"final-project/repositories/user"
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
	repo := NewServiceRepository(db)

	mockUser := SeederUser.UserSeeder()
	mockService := SeederService.ServiceSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Create(mockService)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		res, err := repo.Create(mockService)
		assert.Nil(t, err)
		assert.Equal(t, mockService.Title, res.Title)
	})
}

func TestGet(t *testing.T) {
	Migrator()
	userRepo := user.NewUserRepository(db)
	repo := NewServiceRepository(db)

	mockUser := SeederUser.UserSeeder()
	mockService := SeederService.ServiceSeeder()

	t.Run("negative", func(t *testing.T) {
		res, err := repo.Get()
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		repo.Create(mockService)
		res, err := repo.Get()
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
}

func TestGetDetails(t *testing.T) {
	Migrator()
	userRepo := user.NewUserRepository(db)
	repo := NewServiceRepository(db)

	mockUser := SeederUser.UserSeeder()
	mockService := SeederService.ServiceSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.GetDetails(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		repo.Create(mockService)
		_, err := repo.GetDetails(1)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	Migrator()
	userRepo := user.NewUserRepository(db)
	repo := NewServiceRepository(db)

	mockUser := SeederUser.UserSeeder()
	mockService := SeederService.ServiceSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Update(mockService)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		repo.Create(mockService)

		mockService2 := SeederService.ServiceSeeder()
		mockService2.ID = 1
		mockService2.Title = "service 1"
		_, err := repo.Update(mockService2)
		assert.Nil(t, err)
	})
}

func TestUpdateImage(t *testing.T) {
	Migrator()
	userRepo := user.NewUserRepository(db)
	repo := NewServiceRepository(db)

	mockUser := SeederUser.UserSeeder()
	mockService := SeederService.ServiceSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.UpdateImage(1, "")
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		repo.Create(mockService)

		_, err := repo.UpdateImage(1, "")
		assert.Nil(t, err)
	})
}

func TestDelete(t *testing.T) {
	Migrator()
	userRepo := user.NewUserRepository(db)
	repo := NewServiceRepository(db)

	mockUser := SeederUser.UserSeeder()
	mockService := SeederService.ServiceSeeder()

	t.Run("negative", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		userRepo.Create(mockUser)
		repo.Create(mockService)
		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
