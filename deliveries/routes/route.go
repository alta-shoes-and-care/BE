package routes

import (
	"final-project/deliveries/controllers/auth"
	"final-project/deliveries/controllers/order"
	paymentmethod "final-project/deliveries/controllers/payment-method"
	"final-project/deliveries/controllers/review"
	"final-project/deliveries/controllers/service"
	"final-project/deliveries/controllers/user"
	"final-project/deliveries/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *auth.AuthController, uc *user.UserController, sc *service.ServiceController, pmc *paymentmethod.PaymentMethodController, oc *order.OrderController, rc *review.ReviewController) {
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
	uj.GET("", uc.GetAllUsers())
	uj.GET("/:id", uc.GetByID())
	uj.GET("/me", uc.Get())

	// Payment Method Route
	pm := e.Group("/payments")
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
	sj.DELETE("/:id", sc.Delete())

	// Order route
	o := e.Group("/orders")
	o.Use(middlewares.JWTMiddleware())
	o.POST("", oc.Create())
	o.GET("", oc.Get())
	o.GET("/me", oc.GetByUserID())
	o.GET("/:id", oc.GetByID())
	o.PUT("/check-payment/:id", oc.CheckPaymentStatus())
	o.PUT("/accept/:id", oc.SetAccepted())
	o.PUT("/reject/:id", oc.SetRejected())
	o.PUT("/process/:id", oc.SetOnProcess())
	o.PUT("/deliver/:id", oc.SetDelivering())
	o.PUT("/cancel/:id", oc.SetCancel())
	o.PUT("/done/:id", oc.SetDone())
	o.PUT("/refund/:id", oc.SetRefund())

	// Review Method Route
	r := e.Group("/reviews")
	r.GET("", rc.Get())
	rj := r.Group("/jwt")
	rj.Use(middlewares.JWTMiddleware())
	rj.POST("", rc.Insert())
}
