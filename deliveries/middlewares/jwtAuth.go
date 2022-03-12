package middlewares

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(ID uint, isAdmin bool) (string, error) {
	if ID < 1 {
		return "", errors.New("user_id tidak valid")
	}
	data := jwt.MapClaims{}
	data["id"] = ID
	data["isAdmin"] = isAdmin
	data["expired"] = time.Now().Add(time.Hour * 1).Unix()
	data["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ExtractTokenUserID(e echo.Context) uint {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		data := user.Claims.(jwt.MapClaims)
		id := uint(data["id"].(float64))
		return id
	}
	return 0
}

func ExtractTokenIsAdmin(e echo.Context) bool {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		data := user.Claims.(jwt.MapClaims)
		isAdmin := data["isAdmin"].(bool)
		return isAdmin
	}
	return false
}

func ExtractTokenIsAlive(e echo.Context) bool {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		data := user.Claims.(jwt.MapClaims)
		expiredTime := time.Unix(data["expired"].(int64), 0)

		if remainder := time.Until(expiredTime); remainder > 0 {
			return true
		}
	}
	return false
}
