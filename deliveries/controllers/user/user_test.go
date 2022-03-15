package user

import (
	"bytes"
	"encoding/json"
	"final-project/deliveries/controllers/auth"
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	MockUser "final-project/deliveries/mocks/user"
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
	rootPath = "/users"
	jwtPath  = "/jwt"
)

func TestCreate(t *testing.T) {
	t.Run("fail to bind json", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateUser{
			Name:  "",
			Email: "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		userController := NewUserController(&MockUser.MockUserRepository{})
		userController.Create()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Nil(t, response.Data)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("fail to validate", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateUser{
			Name:     "a",
			Email:    "b",
			Password: "a",
			IsAdmin:  true,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		userController := NewUserController(&MockUser.MockUserRepository{})
		userController.Create()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Nil(t, response.Data)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("fail to create", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateUser{
			Name:     "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
			IsAdmin:  true,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		userController := NewUserController(&MockUser.MockFalseUserRepository{})
		userController.Create()(context)
		context.SetPath(rootPath)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Nil(t, response.Data)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to create", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateUser{
			Name:     "ucup",
			Email:    "ucup@ucup.com",
			Password: "ucup123",
			IsAdmin:  true,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		userController := NewUserController(&MockUser.MockUserRepository{})
		userController.Create()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.NotNil(t, response.Data)
		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGet(t *testing.T) {
	var jwtToken string

	if err := godotenv.Load(".env"); err != nil {
		log.Info("tidak dapat memuat env file", err)
	}

	t.Run("test login", func(t *testing.T) {
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
		jwtToken = dataMap["token"].(string)

		assert.NotEmpty(t, jwtToken)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("fail to get", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockFalseUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.Get())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to get", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.Get())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestGetAllUsers(t *testing.T) {
	var jwtTokenUser, jwtTokenAdmin string

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

	t.Run("login admin", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		authControl := auth.NewAuthController(&MockUser.MockAuthAdminRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtTokenAdmin = dataMap["token"].(string)

		assert.NotEmpty(t, jwtTokenAdmin)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("fail admin", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockFalseAdminRepository{})
		if err := middlewares.JWTMiddleware()(userController.Get())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}
