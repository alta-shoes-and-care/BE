package main

import (
	"final-project/configs"
	_AuthController "final-project/deliveries/controllers/auth"
	_OrderController "final-project/deliveries/controllers/order"
	_PMController "final-project/deliveries/controllers/payment-method"
	_ServiceController "final-project/deliveries/controllers/service"
	_UserController "final-project/deliveries/controllers/user"
	"final-project/deliveries/routes"
	_AuthRepo "final-project/repositories/auth"
	_OrderRepo "final-project/repositories/order"
	_PMRepo "final-project/repositories/payment-method"
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

	authRepo := _AuthRepo.NewAuthRepository(db)
	userRepo := _UserRepo.NewUserRepository(db)
	paymentMethodRepo := _PMRepo.NewPaymentMethodRepository(db)
	serviceRepo := _ServiceRepo.NewServiceRepository(db)
	orderRepo := _OrderRepo.NewOrderRepository(db)

	awsSess := awss3.InitS3(config.S3_KEY, config.S3_SECRET, config.S3_REGION)

	ac := _AuthController.NewAuthController(authRepo)
	uc := _UserController.NewUserController(userRepo)
	pmc := _PMController.NewPaymentMethodController(paymentMethodRepo)
	sc := _ServiceController.NewServiceController(serviceRepo, config, awsSess)
	oc := _OrderController.NewOrderController(orderRepo)

	e := echo.New()

	routes.RegisterPaths(e, ac, uc, sc, pmc, oc)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
