package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-ddd-bank/domain/dto"
	domain "github.com/go-ddd-bank/domain/model"
	errors "github.com/go-ddd-bank/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) CreateUser(user *domain.User) (*dto.UserDTO, *errors.Errors) {
	args := m.Called(user)
	return args.Get(0).(*dto.UserDTO), nil
}

func (m *MockUserService) GetUserByEmail(user *domain.User) (*dto.UserDTO, *errors.Errors) {
	args := m.Called(user)
	return args.Get(0).(*dto.UserDTO), nil
}

func (m *MockUserService) GetUserByID(user *domain.User) (*dto.UserDTO, *errors.Errors) {
	args := m.Called(user)
	return args.Get(0).(*dto.UserDTO), nil
}

func setupTest() (*gin.Engine, *MockUserService, *UserHandler) {
	gin.SetMode(gin.TestMode)

	mockUserService := new(MockUserService)
	mockUserHandler := NewUserHandler(mockUserService)

	router := gin.Default()

	return router, mockUserService, mockUserHandler
}

func TestRegisterUser(t *testing.T) {
	router, mockUserService, mockUserHandler := setupTest()
	router.POST("/api/auth/register", mockUserHandler.RegisterUser)

	t.Run("it should register a user", func(t *testing.T) {
		mockUser := &domain.User{
			ID:        1,
			FirstName: "Ziad",
			LastName:  "Elshimy",
			Password:  "asdasdasd",
			Email:     "Ziadshimy7@gmail.com",
			Phone:     "+79122390087",
		}

		mockReturnedUser := &dto.UserDTO{
			ID:        1,
			FirstName: "Ziad",
			LastName:  "Elshimy",
			Email:     "Ziadshimy7@gmail.com",
			Phone:     "+79122390087",
		}

		mockUserService.On("CreateUser", mockUser).Return(mockReturnedUser, nil)
		body, _ := json.Marshal(mockUser)
		request, _ := http.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer(body))
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		mockUserService.AssertExpectations(t)
	})

	t.Run("it should handle invalid JSON body", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/api/auth/register", bytes.NewBuffer([]byte(`invalid json`)))
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}

func TestLogin(t *testing.T) {
	router, mockUserService, mockUserHandler := setupTest()
	router.POST("/api/auth/login", mockUserHandler.Login)

	t.Run("it should get the user by email, without the password", func(t *testing.T) {
		mockUserEmail := &domain.User{
			Email: "Ziadshimy7@gmail.com",
		}

		mockUser := &dto.UserDTO{
			ID:        1,
			FirstName: "Ziad",
			LastName:  "Elshimy",
			Email:     "Ziadshimy7@gmail.com",
			Phone:     "+79122390087",
		}

		mockUserService.On("GetUserByEmail", mockUserEmail).Return(mockUser)
		body, _ := json.Marshal(mockUserEmail)
		request, _ := http.NewRequest(http.MethodPost, "/api/auth/login", bytes.NewBuffer(body))
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		mockUserService.AssertExpectations(t)
	})
}
