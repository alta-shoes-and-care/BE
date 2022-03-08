package service

import (
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	_ServiceRepo "final-project/repositories/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ServiceController struct {
	repo _ServiceRepo.Service
}

func NewServiceController(repository _ServiceRepo.Service) *ServiceController {
	return &ServiceController{
		repo: repository,
	}
}

func (ctl *ServiceController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newService RequestCreate
		var image string
		userID := middlewares.ExtractTokenUserID(c)

		if err := c.Bind(&newService); err != nil || strings.TrimSpace(newService.Title) == "" || strings.TrimSpace(newService.Description) == "" || newService.Price == 0 {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai, title, description, atau price tidak boleh kosong"))
		}

		res, err := ctl.repo.Create(newService.ToEntityService(image, userID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan Service baru", ToResponseCreate(res)))
	}
}

func (ctl *ServiceController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ctl.repo.Get()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses mendapatkan semua service", ToResponseGet(res)))
	}
}

func (ctl *ServiceController) GetDetails() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))

		res, err := ctl.repo.GetDetails(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses mendapatkan detail service", ToResponseGetDetails(res)))
	}
}

func (ctl *ServiceController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		isAdmin := middlewares.ExtractTokenIsAdmin(c)
		if !isAdmin {
			return c.JSON(http.StatusUnauthorized, common.UnAuthorized("missing or malformed JWT"))
		}

		var updateService RequestUpdate
		var image string
		if err := c.Bind(&updateService); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari user tidak sesuai"))
		}

		res, err := ctl.repo.Update(updateService.ToEntityService(image))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses memperbarui data service", ToResponseUpdate(res)))
	}
}

func (ctl *ServiceController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ID, _ := strconv.Atoi(c.Param("id"))

		err := ctl.repo.Delete(uint(ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menghapus service", err))
	}
}
