package routes

import (
	"final-project/deliveries/controllers/auth"
	paymentmethod "final-project/deliveries/controllers/payment-method"
	"final-project/deliveries/controllers/service"
	"final-project/deliveries/controllers/user"
	"final-project/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *auth.AuthController, uc *user.UserController, sc *service.ServiceController, pmc *paymentmethod.PaymentMethodController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Auth Route
	a := e.Group("/login")
	a.POST("", ac.Login())

	// User Route
	u := e.Group("/users")
	u.POST("", uc.Create())
	uj := u.Group("/jwt")
	uj.Use(middlewares.JWTMiddleware())
	uj.GET("/me", uc.Get())
	uj.PUT("/me", uc.Update())
	uj.DELETE("/me", uc.Delete())

	// Payment Method Route
	pm := e.Group("/payment-methods")
	pm.Use(middlewares.JWTMiddleware())
	pm.POST("", pmc.Create())
	pm.GET("", pmc.Get())
	pm.DELETE("/:id", pmc.Delete())

	// Service Route
	s := e.Group("/services")
	s.GET("", sc.Get())
	s.GET("/:id", sc.GetDetails())
	sj := s.Group("/jwt")
	sj.Use(middlewares.JWTMiddleware())
	sj.POST("", sc.Create())
	sj.PUT("", sc.Update())
	sj.PUT("/:id", sc.UpdateImage())
	sj.DELETE("/:id", sc.Delete())
}
