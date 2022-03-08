package paymentmethod

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	_PMRepo "final-project/repositories/payment-method"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentMethodController struct {
	repo _PMRepo.PaymentMethod
}

func NewPaymentMethodController(repository _PMRepo.PaymentMethod) *PaymentMethodController {
	return &PaymentMethodController{
		repo: repository,
	}
}

func (ctl *PaymentMethodController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		NewPaymentMethod := RequestCreatePaymentMethod{}

		if err := c.Bind(&NewPaymentMethod); err != nil || NewPaymentMethod.Name == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, nama payment method tidak boleh kosong"))
		}

		res, err := ctl.repo.Create(NewPaymentMethod.ToEntityPaymentMethod())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan payment method baru", ToResponseCreatePaymentMethod(res)))
	}
}

func (ctl *PaymentMethodController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}
		res, err := ctl.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua payment method", ToResponseGetPaymentMethod(res)))
	}
}

func (ctl *PaymentMethodController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))

		err := ctl.repo.Delete(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menghapus payment method", err))
	}
}
