package auth

import (
	"bytes"
	"encoding/json"
	"final-project/deliveries/controllers/common"
	MockAuth "final-project/deliveries/mocks/auth"
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
	rootPath = "/login"
)

func TestLogin(t *testing.T) {
	if err := godotenv.Load(".env"); err != nil {
		log.Info("tidak dapat memuat env file", err)
	}

	t.Run("bind error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestLogin{
			Email: "ucup@ucup.com",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		authControl := NewAuthController(&MockAuth.MockAuthRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		authControl := NewAuthController(&MockAuth.MockFalseAuthRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("not acceptable", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		authControl := NewAuthController(&MockAuth.MockFalseAuthRepositoryNotAcceptable{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusNotAcceptable, response.Code)
	})

	t.Run("successful login", func(t *testing.T) {
		var jwtToken string

		requestBody, _ := json.Marshal(RequestLogin{
			Email:    "ucup@ucup.com",
			Password: "ucup123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		authControl := NewAuthController(&MockAuth.MockAuthRepository{})
		authControl.Login()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		dataMap := response.Data.(map[string]interface{})
		jwtToken = dataMap["token"].(string)

		assert.NotEmpty(t, jwtToken)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}
