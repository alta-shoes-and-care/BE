package routes

import (
	"final-project/deliveries/controllers/auth"
	"final-project/deliveries/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *auth.AuthController, uc *user.UserController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))

	a := e.Group("/login")
	a.POST("", ac.Login())

	u := e.Group("/users")
	u.POST("", uc.Create())
}
