package order

import (
	"bytes"
	"encoding/json"
	"final-project/deliveries/controllers/auth"
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	MockOrder "final-project/deliveries/mocks/order"
	MockUser "final-project/deliveries/mocks/user"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var (
	e        = echo.New()
	rootPath = "/orders"
)

func TestCreate(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("user error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:       1,
			Qty:             1,
			Total:           10000,
			PaymentMethodID: 1,
			Date:            "2022-03-18",
			Address:         "Jl. Soedirman",
			City:            "Surabaya",
			Phone:           "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("bind error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         0,
			Qty:               0,
			Total:             0,
			PaymentMethodID:   0,
			PaymentMethodName: "",
			Date:              "",
			Address:           "",
			City:              "",
			Phone:             "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("last order id and midtrans link error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "BCA Klikpay",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockFalseMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("last order id and midtrans link error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "BCA Klikpay",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockFalseMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("get payment url bca klikpay error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "BCA Klikpay",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockFalseRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("get payment url cimb clicks error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "CIMB Clicks",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockFalseRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("get payment url danamon online error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "Danamon Online",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockFalseRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("get payment url gopay error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "Gopay",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockFalseActionsMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("get payment other error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "Other",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockFalseRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("create order error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "BCA Klikpay",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed bca klikpay", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "BCA Klikpay",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("succeed cimb clicks", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "CIMB Clicks",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("succeed danamon online", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "Danamon Online",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("succeed gopay", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateOrder{
			ServiceID:         1,
			Qty:               1,
			Total:             10000,
			PaymentMethodID:   1,
			PaymentMethodName: "Gopay",
			Date:              "2022-03-18",
			Address:           "Jl. Soedirman",
			City:              "Surabaya",
			Phone:             "080000000000",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueActionsMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGet(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("admin error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestGetByUserID(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/me", rootPath))

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.GetByUserID())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/me", rootPath))

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.GetByUserID())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestGetByID(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("as user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.GetByID())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("user is successful", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.GetByID())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("as admin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.GetByID())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("admin is successful", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.GetByID())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestCheckPaymentStatus(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("internal midtrans error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockFalseMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("successful status pending", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("internal error settlement", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockSettlementMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("successful settlement", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockSettlementMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("internal error cancel", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockCancelMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("successful cancel", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockCancelMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("internal error deny", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockDenyMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("successful deny", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockDenyMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("successful accept", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockAcceptMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("successful challenge", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/check-payment", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockChallengeMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.CheckPaymentStatus())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestSetAccepted(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("admin error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/accept", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetAccepted())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/accept", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetAccepted())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/accept", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetAccepted())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestSetRejected(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("admin error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/reject", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetRejected())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/reject", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetRejected())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/reject", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetRejected())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestSetOnProcess(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("admin error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/process", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetOnProcess())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/process", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetOnProcess())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/process", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetOnProcess())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestSetDelivering(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("admin error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/deliver", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetDelivering())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/deliver", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetDelivering())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/deliver", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetDelivering())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestSetCancel(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("user internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/cancel", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetCancel())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("user successful", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/cancel", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetCancel())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("admin internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/cancel", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetCancel())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("admin successful", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/cancel", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetCancel())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestSetDone(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/done", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetDone())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/done", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetDone())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestSetRefund(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("admin error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/refund", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetRefund())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/refund", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockFalseOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetRefund())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/refund", rootPath))
		context.SetParamNames("id")
		context.SetParamValues("1")

		serviceController := NewOrderController(&MockOrder.MockTrueOrderRepository{}, &MockOrder.MockTrueRedirectURLMidtrans{})
		if err := middlewares.JWTMiddleware()(serviceController.SetRefund())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}
