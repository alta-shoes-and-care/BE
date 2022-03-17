package user

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	"final-project/deliveries/validators"
	_UserRepo "final-project/repositories/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo _UserRepo.User
}

func NewUserController(repository _UserRepo.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (ctl *UserController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser RequestCreateUser

		if err := c.Bind(&newUser); err != nil || strings.TrimSpace(newUser.Name) == "" || strings.TrimSpace(newUser.Email) == "" || strings.TrimSpace(newUser.Password) == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, nama, email atau password tidak boleh kosong"))
		}

		if err := validators.ValidateCreateUser(newUser.Name, newUser.Email, newUser.Password); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(err.Error()))
		}

		res, err := ctl.repo.Create(newUser.ToEntityUser())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan user baru", ToResponseCreateUser(res)))
	}
}

func (ctl *UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.Get(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan data user", ToResponseGetUser(res)))
	}
}

func (ctl *UserController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		userID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.GetByID(uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan data user", ToResponseGetUser(res)))
	}
}

func (ctl *UserController) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		res, err := ctl.repo.GetAllUsers()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua user", ToResponseGetUsers(res)))
	}
}
