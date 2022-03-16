package review

import (
	"bytes"
	"encoding/json"
	"final-project/deliveries/controllers/auth"
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	MockUser "final-project/deliveries/mocks/user"
	MockReview "final-project/deliveries/mocks/review"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var (
	e        = echo.New()
	rootPath = "/reviews"
)

func TestInsert(t *testing.T) {
	var jwtTokenUser string

	if err := godotenv.Load(".env"); err != nil {
		log.Info("tidak dapat memuat env file", err)
	}

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("fail to bind json", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestInsertReview{
			ServiceID: 0,
			OrderID: 0,
			Rating: 0,
			Review: "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		reviewController := NewReviewController(&MockReview.MockReviewRepository{})
		if err := middlewares.JWTMiddleware()(reviewController.Insert())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Nil(t, response.Data)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("fail to create", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestInsertReview{
			ServiceID: 1,
			OrderID: 1,
			Rating: 5,
			Review: "Bagus",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		reviewController := NewReviewController(&MockReview.MockFalseReviewRepository{})
		if err := middlewares.JWTMiddleware()(reviewController.Insert())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Nil(t, response.Data)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("success to create", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestInsertReview{
			ServiceID: 1,
			OrderID: 1,
			Rating: 5,
			Review: "Bagus",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))
		
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		reviewController := NewReviewController(&MockReview.MockReviewRepository{})
		if err := middlewares.JWTMiddleware()(reviewController.Insert())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.NotNil(t, response.Data)
		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGet(t *testing.T) {
	t.Run("fail to get all reviews", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		reviewController := NewReviewController(&MockReview.MockFalseReviewRepository{})
		reviewController.Get()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("success to get all reviews", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		reviewController := NewReviewController(&MockReview.MockReviewRepository{})
		reviewController.Get()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestUpdate(t *testing.T) {
	var jwtTokenUser string

	if err := godotenv.Load(".env"); err != nil {
		log.Info("tidak dapat memuat env file", err)
	}

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("fail to bind", func(t *testing.T) {
		failStruct := struct {
			Rating string `json:"rating"`
		}{"5"}
		requestBody, _ := json.Marshal(failStruct)

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/1", rootPath))

		reviewController := NewReviewController(&MockReview.MockReviewRepository{})
		if err := middlewares.JWTMiddleware()(reviewController.Update())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("fail to update", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestUpdateReview{})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/1", rootPath))

		reviewController := NewReviewController(&MockReview.MockFalseReviewRepository{})
		if err := middlewares.JWTMiddleware()(reviewController.Update())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to update", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestUpdateReview{
			Rating: 5,
			Review: "Bagus",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/1", rootPath))

		reviewController := NewReviewController(&MockReview.MockReviewRepository{})
		if err := middlewares.JWTMiddleware()(reviewController.Update())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestDelete(t *testing.T) {
	var jwtTokenUser string

	if err := godotenv.Load(".env"); err != nil {
		log.Info("tidak dapat memuat env file", err)
	}

	t.Run("login user", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		authControl := auth.NewAuthController(&MockUser.MockAuthUserRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenUser = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenUser)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("fail to delete", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/1", rootPath))

		reviewController := NewReviewController(&MockReview.MockFalseReviewRepository{})
		if err := middlewares.JWTMiddleware()(reviewController.Delete())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to delete", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v/1", rootPath))

		reviewController := NewReviewController(&MockReview.MockReviewRepository{})
		if err := middlewares.JWTMiddleware()(reviewController.Delete())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}
