package order

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	_OrderRepo "final-project/repositories/order"
	midtranspay "final-project/services/midtrans-pay"
	"net/http"
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
	layoutISO = "2006-01-02 WIB"
)

func (ctl *OrderController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
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
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan Order baru", res))
	}
}
