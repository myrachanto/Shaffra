package users

import ( // Adjust the import path to your project structure

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the UserServiceInterface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(user *User) (*User, error) {
	args := m.Called(user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserService) GetOne(id string) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserService) GetAll() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockUserService) Update(id string, user *User) (*User, error) {
	args := m.Called(id, user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserService) Delete(id string) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}
func TestCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	router := gin.Default()
	router.POST("/users/create", controller.Create)

	mockUser := &User{
		Name:  "John",
		Email: "john@example.com",
		Age:   23,
	}

	mockService.On("Create", mockUser).Return(mockUser, nil)

	form := `name=John&email=john@example.com&age=23`
	req, _ := http.NewRequest(http.MethodPost, "/users/create", strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Unmarshal the response and only compare the necessary fields
	var actual map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &actual)

	user := actual["user"].(map[string]interface{})
	assert.Equal(t, "John", user["name"])
	assert.Equal(t, "john@example.com", user["email"])
	assert.Equal(t, float64(23), user["age"]) // age is returned as float64

	mockService.AssertExpectations(t)
}

func TestGetOne(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	router := gin.Default()
	router.GET("/users/:id", controller.GetOne)

	mockUser := &User{
		Name:  "John",
		Email: "john@example.com",
		Age:   23,
	}

	mockService.On("GetOne", "1").Return(mockUser, nil)

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Unmarshal the response and only compare the necessary fields
	var actual map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &actual)

	user := actual["user"].(map[string]interface{})
	assert.Equal(t, "John", user["name"])
	assert.Equal(t, "john@example.com", user["email"])
	assert.Equal(t, float64(23), user["age"]) // age is returned as float64

	mockService.AssertExpectations(t)
}

func TestGetAll(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	router := gin.Default()
	router.GET("/api/users", controller.GetAll)

	mockUsers := []User{
		{Name: "John", Email: "john@example.com", Age: 23},
		{Name: "Jane", Email: "jane@example.com", Age: 30},
	}

	mockService.On("GetAll").Return(mockUsers, nil)

	req, _ := http.NewRequest(http.MethodGet, "/api/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	mockService.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	router := gin.Default()
	router.PUT("/api/users/:id", controller.Update)

	mockUser := &User{
		Name:  "John",
		Email: "john@example.com",
		Age:   23,
	}

	mockService.On("Update", "1", mockUser).Return(mockUser, nil)

	form := `name=John&email=john@example.com&age=23`
	req, _ := http.NewRequest(http.MethodPut, "/api/users/1", strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `{"status":"Updated successifuly"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	mockService.AssertExpectations(t)
}
func TestDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	controller := NewUserController(mockService)

	router := gin.Default()
	router.DELETE("/api/users/:id", controller.Delete)

	mockService.On("Delete", "1").Return("deleted successfully", nil)

	req, _ := http.NewRequest(http.MethodDelete, "/api/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedResponse := `{"status":"deleted successfully"}`
	assert.JSONEq(t, expectedResponse, w.Body.String())

	mockService.AssertExpectations(t)
}
