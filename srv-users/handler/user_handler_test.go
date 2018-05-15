package handler

import (
	"encoding/json"
	"github.com/dfibrinogen/dfibrinogen-api/common"
	"github.com/dfibrinogen/dfibrinogen-api/srv-users/model"
	"github.com/dfibrinogen/dfibrinogen-api/srv-users/repository"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewUserHandler(t *testing.T) {

	mockRepo := new(repository.MockUserRepository)

	e := echo.New()
	g := e.Group("/api/v1")
	NewUserHandler(g, mockRepo)
}

func TestUserHandler_GetDataAll(t *testing.T) {

	mockData := []model.User{{
		ID:        "USER.01",
		Name:      "Test 01",
		Gender:    "P",
		Age:       23,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}}

	mockResponse := common.Response{
		Status:  http.StatusOK,
		Message: common.SUCCESS_GET_DATA_ALL,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockUserRepository)
	mockRepo.On("FetchUserAll").
		Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	h := &userHandler{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.GetDataAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestUserHandler_GetDataAll_Error(t *testing.T) {

}

func TestUserHandler_GetDataByID(t *testing.T) {

	mockData := model.User{
		ID:        "USER.01",
		Name:      "Test 01",
		Gender:    "P",
		Age:       23,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}

	mockResponse := common.Response{
		Status:  http.StatusOK,
		Message: common.SUCCESS_GET_DATA,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockUserRepository)
	mockRepo.On("FetchUserByID", "USER.01").
		Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("USER.01")

	h := &userHandler{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.GetDataByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestUserHandler_GetDataByID_Error(t *testing.T) {

}

func TestUserHandler_CreateData(t *testing.T) {

}

func TestUserHandler_CreateData_Error(t *testing.T) {

}

func TestUserHandler_CreateData_Bind_Error(t *testing.T) {

}

func TestUserHandler_UpdateData(t *testing.T) {

}

func TestUserHandler_UpdateData_Error(t *testing.T) {

}

func TestUserHandler_UpdateData_Bind_Error(t *testing.T) {

}

func TestUserHandler_DeleteData(t *testing.T) {

}

func TestUserHandler_DeleteData_Error(t *testing.T) {

}
