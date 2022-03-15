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

func TestGetByID(t *testing.T) {
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
		context.SetPath(fmt.Sprintf("%v%v/1", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.GetByID())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("fail to get user by id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v/1", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockFalseUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.GetByID())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to get by id", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v/1", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.GetByID())(context); err != nil {
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

		userController := NewUserController(&MockUser.MockUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.GetAllUsers())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("fail to get all users", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockFalseAdminRepository{})
		if err := middlewares.JWTMiddleware()(userController.GetAllUsers())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to get all users", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.GetAllUsers())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestUpdate(t *testing.T) {
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

	t.Run("fail to bind", func(t *testing.T) {
		failStruct := struct {
			Name int `json:"name"`
		}{1}
		requestBody, _ := json.Marshal(failStruct)

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v/me", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.Update())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("fail to validate", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestUpdateUser{
			Name:     "a",
			Email:    "b",
			Password: "x",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v/me", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.Update())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("fail to update", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestUpdateUser{})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v/me", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockFalseUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.Update())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to update", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestUpdateUser{
			Name: "Ucup Updated",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v/me", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.Update())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestDelete(t *testing.T) {
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

	t.Run("fail to delete", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v/me", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockFalseUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.Delete())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to delete", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v/me", rootPath, jwtPath))

		userController := NewUserController(&MockUser.MockUserRepository{})
		if err := middlewares.JWTMiddleware()(userController.Delete())(context); err != nil {
			log.Fatal(err)
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}
