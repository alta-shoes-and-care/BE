package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BodyLimiter() echo.MiddlewareFunc {
	return middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{
		Limit: "500KB",
	})
}
