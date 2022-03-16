package service

import (
	"bytes"
	"encoding/json"
	"final-project/configs"
	"final-project/deliveries/controllers/auth"
	"final-project/deliveries/controllers/common"
	"final-project/deliveries/middlewares"
	MockService "final-project/deliveries/mocks/service"
	MockUser "final-project/deliveries/mocks/user"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var (
	e        = echo.New()
	rootPath = "/services"
	jwtPath  = "/jwt"
	config   = configs.GetConfig(true)
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

	t.Run("admin error", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenUser))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		serviceController := NewServiceController(&MockService.MockServiceTrueRepository{}, config, &MockService.MockAWSStructTrue{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusUnauthorized, response.Code)
	})

	t.Run("bind error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreate{
			Title: "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		serviceController := NewServiceController(&MockService.MockServiceTrueRepository{}, config, &MockService.MockAWSStructTrue{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("validation error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreate{
			Title:       "asdhajkdhajsdhjkabsdjabdjkabdkasdbjkabsdjkasbdjkasbdjkasdjabsdjabsdjabsd",
			Description: "Layanan 1",
			Price:       15000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		serviceController := NewServiceController(&MockService.MockServiceTrueRepository{}, config, &MockService.MockAWSStructTrue{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("file error", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreate{
			Title:       "Service 1",
			Description: "Layanan 1",
			Price:       15000,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))
		context.FormFile("file")

		serviceController := NewServiceController(&MockService.MockServiceTrueRepository{}, config, &MockService.MockAWSStructTrue{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("upload error", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		writer.WriteField("title", "Service 1")
		writer.WriteField("description", "Layanan 1")
		writer.WriteField("price", "15000")

		part, err := writer.CreateFormFile("file", "../../../images/shoes-service-station.png")
		if err != nil {
			log.Fatal(err)
			return
		}

		part.Write([]byte("file"))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/", body)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		serviceController := NewServiceController(&MockService.MockServiceTrueRepository{}, config, &MockService.MockAWSStructFalse{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("internal server error", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		writer.WriteField("title", "Service 1")
		writer.WriteField("description", "Layanan 1")
		writer.WriteField("price", "15000")

		part, err := writer.CreateFormFile("file", "../../../images/shoes-service-station.png")
		if err != nil {
			log.Fatal(err)
			return
		}

		part.Write([]byte("file"))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/", body)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		serviceController := NewServiceController(&MockService.MockServiceFalseRepository{}, config, &MockService.MockAWSStructTrue{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("success", func(t *testing.T) {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		writer.WriteField("title", "Service 1")
		writer.WriteField("description", "Layanan 1")
		writer.WriteField("price", "15000")

		part, err := writer.CreateFormFile("file", "../../../images/shoes-service-station.png")
		if err != nil {
			log.Fatal(err)
			return
		}

		part.Write([]byte("file"))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/", body)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtTokenAdmin))

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v%v", rootPath, jwtPath))

		serviceController := NewServiceController(&MockService.MockServiceTrueRepository{}, config, &MockService.MockAWSStructTrue{})
		if err := middlewares.JWTMiddleware()(serviceController.Create())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}
