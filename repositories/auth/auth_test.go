package auth

// import (
// 	"group-project2/configs"
// 	B "group-project2/entities/book"
// 	I "group-project2/entities/image"
// 	PM "group-project2/entities/payment-method"
// 	Rat "group-project2/entities/rating"
// 	R "group-project2/entities/room"
// 	U "group-project2/entities/user"
// 	"group-project2/repositories/user"
// 	"group-project2/utils"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	config = configs.GetConfig(true)
// 	db     = utils.InitDB(config)
// )

// func TestLogin(t *testing.T) {
// 	db.Migrator().DropTable(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})
// 	db.AutoMigrate(U.Users{}, PM.PaymentMethods{}, R.Rooms{}, Rat.Ratings{}, I.Images{}, B.Books{})

// 	repo := New(db)
// 	UR := user.New(db)

// 	t.Run("fail to login", func(t *testing.T) {
// 		mockUser := U.Users{
// 			Email:    "ucup@ucup.com",
// 			Password: "ucup123",
// 		}

// 		_, err := repo.Login(mockUser.Email, mockUser.Password)
// 		assert.NotNil(t, err)
// 	})

// 	t.Run("success to login", func(t *testing.T) {
// 		mockUser := U.Users{
// 			Name:     "Ucup",
// 			Email:    "ucup@ucup.com",
// 			Password: "ucup123",
// 		}
// 		UR.Insert(mockUser)

// 		_, err := repo.Login(mockUser.Email, mockUser.Password)
// 		assert.Nil(t, err)
// 	})
// }
