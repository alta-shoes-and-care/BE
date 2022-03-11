package user

import (
	"final-project/configs"
	O "final-project/entities/order"
	PM "final-project/entities/payment-method"
	R "final-project/entities/review"
	S "final-project/entities/service"
	U "final-project/entities/user"
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
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("positive", func(t *testing.T) {
		res, err := repo.Create(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Create(mockUser)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.Get(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		res, err := repo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})
}

func TestGetByID(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.GetByID(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		res, err := repo.GetByID(1)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})
}

func TestGetUsers(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("negative", func(t *testing.T) {
		_, err := repo.GetUsers()
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		res, err := repo.GetUsers()
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res[0].Name)
	})
}

func TestUpdate(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		mockUser2 := SeederUser.UserSeeder()
		mockUser2.ID = 1
		mockUser2.Name = "ucup_Updated"

		res, err := repo.Update(mockUser2)
		assert.Nil(t, err)
		assert.Equal(t, mockUser2.Name, res.Name)
	})

	t.Run("negative", func(t *testing.T) {
		mockUser2 := SeederUser.UserSeeder()

		_, err := repo.Update(mockUser2)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	Migrator()
	repo := NewUserRepository(db)

	mockUser := SeederUser.UserSeeder()

	t.Run("negative", func(t *testing.T) {
		err := repo.Delete(1)
		assert.NotNil(t, err)
	})

	t.Run("positive", func(t *testing.T) {
		repo.Create(mockUser)

		err := repo.Delete(1)
		assert.Nil(t, err)
	})
}
