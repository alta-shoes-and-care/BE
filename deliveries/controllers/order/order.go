package order

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	_OrderRepo "final-project/repositories/order"
	midtranspay "final-project/services/midtrans-pay"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	repo _OrderRepo.Order
}

func NewOrderController(repository _OrderRepo.Order) *OrderController {
	return &OrderController{
		repo: repository,
	}
}

var (
	midtransClient = midtranspay.InitConnection()
)

const (
	layoutISO = "2006-01-02"
)

func (ctl *OrderController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		var newOrder RequestCreateOrder
		if err := c.Bind(&newOrder); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, service_id, payment_method_id, time, address, city, atau phone tidak boleh kosong"))
		}

		date, _ := time.Parse(layoutISO, newOrder.Date)
		userID := middlewares.ExtractTokenUserID(c)
		order, err := ctl.repo.Create(newOrder.ToEntityOrder(date, userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}

		midtransCharge := midtranspay.CreateTransaction(midtransClient, order.ID, newOrder.Total)
		url := midtransCharge.RedirectURL
		res, err := ctl.repo.InsertUrl(order.ID, url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan Order baru", ToResponseOrder(res)))
	}
}

func (ctl *OrderController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		res, err := ctl.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua order", ToResponseOrderArr(res)))
	}
}

func (ctl *OrderController) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		userID := middlewares.ExtractTokenUserID(c)
		res, err := ctl.repo.GetByUserID(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua order berdasarkan user id", ToResponseOrderArr(res)))
	}
}

func (ctl *OrderController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ctl.repo.GetByID(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan detail order", ToResponseOrder(res)))
	}
}

func (ctl *OrderController) CheckPaymentStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := midtranspay.CheckTransaction(midtransClient, uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}

		if res == "status pending" {
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "pembayaran tertunda", nil))
		} else if res == "status settlement" {
			res, err := ctl.repo.SetPaid(uint(ID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
			}
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses menjadikan status pembayaran menjadi paid", ToResponseOrder(res)))
		} else if res == "status cancel" {
			res, err := ctl.repo.SetCancel(uint(ID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
			}
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi cancel", ToResponseOrder(res)))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "status transaksi:", res))
	}
}

func (ctl *OrderController) SetAccepted() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ctl.repo.SetAccepted(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi accepted", ToResponseOrder(res)))
	}
}

func (ctl *OrderController) SetRejected() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ctl.repo.SetRejected(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi rejected", ToResponseOrder(res)))
	}
}

func (ctl *OrderController) SetOnProcess() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ctl.repo.SetOnProcess(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi on process", ToResponseOrder(res)))
	}
}

func (ctl *OrderController) SetDelivering() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ctl.repo.SetDelivering(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi delivering", ToResponseOrder(res)))
	}
}

func (ctl *OrderController) SetCancel() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ctl.repo.SetCancel(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi cancel", res))
	}
}

func (ctl *OrderController) SetDone() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAlive := middlewares.ExtractTokenIsAlive(c)
		if !isAlive {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("JWT token is expired"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))
		res, err := ctl.repo.SetDone(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi done", ToResponseOrder(res)))
	}
}
