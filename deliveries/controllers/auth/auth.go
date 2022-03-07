package auth

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	_AuthRepo "final-project/repositories/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo _AuthRepo.Auth
}

func NewAuthController(repository _AuthRepo.Auth) *AuthController {
	return &AuthController{
		repo: repository,
	}
}

func (ctl *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil || loginFormat.Email == "" || loginFormat.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, email atau password tidak boleh kosong"))
		}

		checkedUser, err := ctl.repo.Login(loginFormat.Email, loginFormat.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}

		IsAdmin := checkedUser.IsAdmin
		tokenID, err := middlewares.GenerateToken(checkedUser.ID, IsAdmin)
		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.NotAcceptable())
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "berhasil masuk, mendapatkan token baru", tokenID))
	}
}
