package user

import (
	"bytes"
	"encoding/json"
	"final-project/deliveries/controllers/auth"
	"final-project/deliveries/controllers/common"
	MockUser "final-project/deliveries/mocks/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	e = echo.New()
)

type LoginDataStruct struct {
	Token   string `json:"token"`
	IsAdmin bool   `json:"is_admin"`
}

func TestCreate(t *testing.T) {
	var jwtToken string

	t.Run("test login", func(t *testing.T) {
		requestBody, _ := json.Marshal(auth.RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		authControl := auth.NewAuthController(&MockUser.MockAuthRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtToken = dataMap["token"].(string)

		assert.NotEmpty(t, jwtToken)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("fail to bind json", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateUser{
			Name:  "",
			Email: "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

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

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Nil(t, response.Data)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("success to create", func(t *testing.T) {
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

		userController := NewUserController(&MockUser.MockUserRepository{})
		userController.Create()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.NotNil(t, response.Data)
		assert.Equal(t, http.StatusCreated, response.Code)
	})
}
