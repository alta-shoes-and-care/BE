package user

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	_UserRepo "final-project/repositories/user"
	"net/http"

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
		NewUser := RequestCreateUser{}

		if err := c.Bind(&NewUser); err != nil || NewUser.Name == "" || NewUser.Email == "" || NewUser.Password == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, nama, email atau password tidak boleh kosong"))
		}

		res, err := ctl.repo.Create(NewUser.ToEntityUser())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan user baru", ToResponseCreateUser(res)))
	}
}

func (ctl *UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.Get(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan data user", ToResponseGetUser(res)))
	}
}

func (ctl *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)
		var UpdatedUser = RequestUpdateUser{}

		if err := c.Bind(&UpdatedUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("terdapat kesalahan input dari client"))
		}

		res, err := ctl.repo.Update(UpdatedUser.ToEntityUser(uint(UserID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses memperbarui data user", ToResponseUpdateUser(res)))
	}
}

func (ctl *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		UserID := middlewares.ExtractTokenUserID(c)

		err := ctl.repo.Delete(uint(UserID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus user", err))
	}
}
