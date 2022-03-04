package main

import (
	"final-project/configs"
	_AuthController "final-project/deliveries/controllers/auth"
	_UserController "final-project/deliveries/controllers/user"
	"final-project/deliveries/routes"
	_AuthRepo "final-project/repositories/auth"
	_UserRepo "final-project/repositories/user"
	"final-project/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig(false)
	db := utils.InitDB(config)

	authRepo := _AuthRepo.New(db)
	userRepo := _UserRepo.NewUserRepository(db)

	// awsSess := awss3.InitS3(config.S3_KEY, config.S3_SECRET, config.S3_REGION)

	ac := _AuthController.New(authRepo)
	uc := _UserController.New(userRepo)

	e := echo.New()

	routes.RegisterPaths(e, ac, uc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
