package main

import (
	"final-project/configs"
	_AuthController "final-project/deliveries/controllers/auth"
	_ServiceController "final-project/deliveries/controllers/service"
	_UserController "final-project/deliveries/controllers/user"
	"final-project/deliveries/routes"
	_AuthRepo "final-project/repositories/auth"
	_ServiceRepo "final-project/repositories/service"
	_UserRepo "final-project/repositories/user"
	awss3 "final-project/services/aws-s3"
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
	serviceRepo := _ServiceRepo.NewServiceRepository(db)

	awsSess := awss3.InitS3(config.S3_KEY, config.S3_SECRET, config.S3_REGION)

	ac := _AuthController.NewAuthController(authRepo)
	uc := _UserController.NewUserController(userRepo)
	sc := _ServiceController.NewServiceController(serviceRepo, config, awsSess)

	e := echo.New()

	routes.RegisterPaths(e, ac, uc, sc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
