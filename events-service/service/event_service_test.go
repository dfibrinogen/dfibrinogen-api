package service

import (
	"encoding/json"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/model"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/repository"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/util"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestEventService_GetDataAll(t *testing.T) {

	mockData := []model.Event{{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}}

	mockResponse := util.Response{
		Status:  http.StatusOK,
		Message: util.SUCCESS_GET_DATA_ALL,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.EventRepo)
	mockRepo.On("FetchEventAll").Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events")

	h := &eventService{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.GetDataAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventService_GetDataByID(t *testing.T) {

	mockData := model.Event{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}

	mockResponse := util.Response{
		Status:  http.StatusOK,
		Message: util.SUCCESS_GET_DATA,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.EventRepo)
	mockRepo.On("FetchEventByID").Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventService{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.GetDataByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventService_CreateData(t *testing.T) {

	mockData := model.Event{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}

	mockJSON, _ := json.Marshal(mockData)

	mockResponse := util.Response{
		Status:  http.StatusCreated,
		Message: util.SUCCESS_CREATE_DATA,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.EventRepo)
	mockRepo.On("CreateEvent", mock.AnythingOfType("model.Event")).Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(mockJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events")

	h := &eventService{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.CreateData(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventService_UpdateData(t *testing.T) {

	mockData := model.Event{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}

	mockJSON, _ := json.Marshal(mockData)

	mockResponse := util.Response{
		Status:  http.StatusOK,
		Message: util.SUCCESS_UPDATE_DATA,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.EventRepo)
	mockRepo.On("UpdateEvent", mock.AnythingOfType("model.Event")).Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/", strings.NewReader(string(mockJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventService{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.UpdateData(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventService_DeleteData(t *testing.T) {

	mockResponse := util.Response{
		Status:  http.StatusOK,
		Message: util.SUCCESS_DELETE_DATA,
		Data:    nil,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.EventRepo)
	mockRepo.On("DeleteEvent", "EVENT.01").Return(nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventService{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.DeleteData(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}
