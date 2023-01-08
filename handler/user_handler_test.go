package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MochamadAkbar/ordent-test/api"
	"github.com/MochamadAkbar/ordent-test/config"
	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/MochamadAkbar/ordent-test/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userHandler = UserHandlerImpl{Usecase: userUsecase}
	userUsecase = &usecase.UserUseCaseMock{Mock: mock.Mock{}}
)

func TestUserHandler_Login(t *testing.T) {
	t.Run("Test User Handler Login Success", func(t *testing.T) {
		var resp api.UserResponse
		user := &entity.User{
			Email:    "test@gmail.com",
			Password: "Password",
		}

		router := config.NewRouter()
		router.Post("/api/v1/authentication/login", userHandler.Login)

		requestBody := strings.NewReader(`{ "email": "test@gmail.com", "password": "Password" }`)
		request := httptest.NewRequest(http.MethodPost,
			"http://localhost:5000/api/v1/authentication/login", requestBody)

		recorder := httptest.NewRecorder()
		userUsecase.Mock.On("Login", user).Return(resp, nil)

		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close() //nolint:errcheck
		assert.Equal(t, "200 OK", response.Status)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	})
}

func TestUserHandler_Register(t *testing.T) {
	t.Run("Test User Handler Register Success", func(t *testing.T) {
		var resp api.UserResponse
		user := &entity.User{
			Name:     "john doe",
			Email:    "test@gmail.com",
			Password: "Password",
		}

		router := config.NewRouter()
		router.Post("/api/v1/authentication/register", userHandler.Register)

		requestBody := strings.NewReader(`{ "name": "john doe", "email": "test@gmail.com", "password": "Password" }`)
		request := httptest.NewRequest(http.MethodPost,
			"http://localhost:5000/api/v1/authentication/register", requestBody)

		recorder := httptest.NewRecorder()
		userUsecase.Mock.On("Register", user).Return(resp, nil)

		router.ServeHTTP(recorder, request)

		response := recorder.Result()
		defer response.Body.Close() //nolint:errcheck
		assert.Equal(t, "201 Created", response.Status)
		assert.Equal(t, http.StatusCreated, response.StatusCode)
	})
}
