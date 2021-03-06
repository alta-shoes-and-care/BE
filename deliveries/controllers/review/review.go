package review

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	_ReviewRepo "final-project/repositories/review"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReviewController struct {
	repo _ReviewRepo.Review
}

func NewReviewController(repository _ReviewRepo.Review) *ReviewController {
	return &ReviewController{
		repo: repository,
	}
}

func (ctl *ReviewController) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := middlewares.ExtractTokenUserID(c)
		NewReview := RequestInsertReview{}

		if err := c.Bind(&NewReview); err != nil || NewReview.Review == "" || NewReview.Rating == 0 || NewReview.ServiceID == 0 || NewReview.OrderID == 0 {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, service_id, order_id, rating, atau review tidak boleh kosong"))
		}

		res, err := ctl.repo.Insert(NewReview.ToEntityReview(uint(userID)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan review baru", res))
	}
}

func (ctl *ReviewController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ctl.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua review", res))
	}
}
