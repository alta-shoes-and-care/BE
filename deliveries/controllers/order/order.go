package order

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	midtranspay "final-project/external/midtrans-pay"
	_OrderRepo "final-project/repositories/order"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	repo           _OrderRepo.Order
	midtransClient midtranspay.MidtransClient
}

func NewOrderController(repository _OrderRepo.Order, client midtranspay.MidtransClient) *OrderController {
	return &OrderController{
		repo:           repository,
		midtransClient: client,
	}
}

const (
	layoutISO = "2006-01-02"
)

// Create new order
func (ctl *OrderController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed jwt"))
		}

		var newOrder RequestCreateOrder
		if err := c.Bind(&newOrder); err != nil || newOrder.ServiceID == 0 || newOrder.Qty == 0 || newOrder.Total == 0 || newOrder.PaymentMethodID == 0 || strings.TrimSpace(newOrder.Date) == "" || strings.TrimSpace(newOrder.Address) == "" || strings.TrimSpace(newOrder.City) == "" || strings.TrimSpace(newOrder.Phone) == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, service_id, qty, total, payment_method_id, date, address, city, atau phone tidak boleh kosong"))
		}

		lastID, err := ctl.repo.GetLastOrderID()
		if err != nil {
			lastID = 0
		}

		userID := middlewares.ExtractTokenUserID(c)
		midtransCharge := ctl.midtransClient.CreateTransaction(userID, lastID+1, newOrder.Total)

		url := midtransCharge.RedirectURL
		if strings.TrimSpace(url) == "" {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError("gagal membuat invoice pembayaran"))
		}

		date, _ := time.Parse(layoutISO, newOrder.Date)

		res, err := ctl.repo.Create(newOrder.ToEntityOrder(date, userID, url))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan Order baru", ToResponseOrder(res)))
	}
}

// Get all orders
func (ctl *OrderController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed jwt"))
		}

		res, err := ctl.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua order", ToResponseOrderArr(res)))
	}
}

// Get all orders by user id
func (ctl *OrderController) GetByUserID() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)

		res, err := ctl.repo.GetByUserID(userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua order berdasarkan user id", ToResponseOrderArr(res)))
	}
}

// Get order by id
func (ctl *OrderController) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			ID, _ := strconv.Atoi(c.Param("id"))
			userID := middlewares.ExtractTokenUserID(c)

			res, err := ctl.repo.GetByIDUser(uint(ID), userID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
			}
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan detail order oleh user", ToResponseOrder(res)))
		} else {
			ID, _ := strconv.Atoi(c.Param("id"))

			res, err := ctl.repo.GetByIDAdmin(uint(ID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
			}
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan detail order oleh admin", ToResponseOrder(res)))
		}
	}
}

// Check payment status of order
func (ctl *OrderController) CheckPaymentStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)
		ID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.midtransClient.CheckTransaction(userID, uint(ID))
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
		} else if res == "status cancel" || res == "status deny" {
			res, err := ctl.repo.SetCancelAdmin(uint(ID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
			}
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi cancel", ToResponseOrder(res)))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "status transaksi:", res))
	}
}

// Set order status accepted
func (ctl *OrderController) SetAccepted() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed jwt"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.SetAccepted(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi accepted", ToResponseOrder(res)))
	}
}

// Set order status rejected
func (ctl *OrderController) SetRejected() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed jwt"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.SetRejected(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi rejected", ToResponseOrder(res)))
	}
}

// Set order status on process
func (ctl *OrderController) SetOnProcess() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed jwt"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.SetOnProcess(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi on process", ToResponseOrder(res)))
	}
}

// Set order status delivering
func (ctl *OrderController) SetDelivering() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed jwt"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.SetDelivering(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi delivering", ToResponseOrder(res)))
	}
}

// Set order status cancel
func (ctl *OrderController) SetCancel() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			ID, _ := strconv.Atoi(c.Param("id"))
			userID := middlewares.ExtractTokenUserID(c)

			res, err := ctl.repo.SetCancelUser(uint(ID), userID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
			}
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi cancel oleh user", ToResponseOrder(res)))
		} else {
			ID, _ := strconv.Atoi(c.Param("id"))

			res, err := ctl.repo.SetCancelAdmin(uint(ID))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
			}
			return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi cancel admin", ToResponseOrder(res)))
		}
	}
}

// Set order status done
func (ctl *OrderController) SetDone() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)
		ID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.SetDone(uint(ID), uint(userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status order menjadi done", ToResponseOrder(res)))
	}
}

// Set order refund status true
func (ctl *OrderController) SetRefund() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed jwt"))
		}

		ID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.SetRefund(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mengubah status refund menjadi true", ToResponseOrder(res)))
	}
}
