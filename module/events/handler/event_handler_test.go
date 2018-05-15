package handler

import (
	"encoding/json"
	"errors"
	"github.com/dfibrinogen/dfibrinogen-api/module"
	"github.com/dfibrinogen/dfibrinogen-api/module/events"
	"github.com/dfibrinogen/dfibrinogen-api/module/events/repository"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewEventHandler(t *testing.T) {

	mockRepo := new(repository.MockEventRepository)

	e := echo.New()
	g := e.Group("/api/v1")
	NewEventHandler(g, mockRepo)
}

func TestEventHandler_GetDataAll(t *testing.T) {

	mockData := []events.Event{{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}}

	mockResponse := module.Response{
		Status:  http.StatusOK,
		Message: module.SUCCESS_GET_DATA_ALL,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("FetchEventAll").
		Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.GetDataAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_GetDataAll_Error(t *testing.T) {

	mockError := module.EMPTY_ERROR

	mockResponse := module.Response{
		Status:  http.StatusNotFound,
		Message: mockError.Error(),
		Data:    nil,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("FetchEventAll").
		Return(nil, mockError).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.Error(t, mockError, h.GetDataAll(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_GetDataByID(t *testing.T) {

	mockData := events.Event{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}

	mockResponse := module.Response{
		Status:  http.StatusOK,
		Message: module.SUCCESS_GET_DATA,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("FetchEventByID", "EVENT.01").
		Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.GetDataByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_GetDataByID_Error(t *testing.T) {

	mockError := module.NOT_FOUND_ERROR

	mockResponse := module.Response{
		Status:  http.StatusNotFound,
		Message: mockError.Error(),
		Data:    nil,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("FetchEventByID", "EVENT.01").
		Return(events.Event{}, mockError).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.Error(t, mockError, h.GetDataByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_CreateData(t *testing.T) {

	mockData := events.Event{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}

	mockJSON, _ := json.Marshal(mockData)

	mockResponse := module.Response{
		Status:  http.StatusCreated,
		Message: module.SUCCESS_CREATE_DATA,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("CreateEvent", mock.AnythingOfType("events.Event")).
		Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(mockJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.CreateData(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_CreateData_Error(t *testing.T) {

	mockError := module.FAILED_SAVE_ERROR

	mockData := events.Event{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}

	mockJSON, _ := json.Marshal(mockData)

	mockResponse := module.Response{
		Status:  http.StatusBadRequest,
		Message: mockError.Error(),
		Data:    nil,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("CreateEvent", mock.AnythingOfType("events.Event")).
		Return(events.Event{}, mockError).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(mockJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.Error(t, mockError, h.CreateData(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_CreateData_Bind_Error(t *testing.T) {

	mockError := errors.New("code=400, message=Unmarshal type error: expected=events.Event, got=array, offset=1")

	mockJSON := `[{"error_1":"01","error_2":"02"}]`

	mockResponse := module.Response{
		Status:  http.StatusBadRequest,
		Message: mockError.Error(),
		Data:    nil,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(mockJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.Error(t, mockError, h.CreateData(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_UpdateData(t *testing.T) {

	mockData := events.Event{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}

	mockJSON, _ := json.Marshal(mockData)

	mockResponse := module.Response{
		Status:  http.StatusOK,
		Message: module.SUCCESS_UPDATE_DATA,
		Data:    mockData,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("UpdateEvent", mock.AnythingOfType("events.Event")).
		Return(mockData, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/", strings.NewReader(string(mockJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.UpdateData(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_UpdateData_Error(t *testing.T) {

	mockError := module.FAILED_UPDATE_ERROR

	mockData := events.Event{
		ID:        "EVENT.01",
		Name:      "Test 01",
		Location:  "Location 01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil}

	mockJSON, _ := json.Marshal(mockData)

	mockResponse := module.Response{
		Status:  http.StatusBadRequest,
		Message: mockError.Error(),
		Data:    nil,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("UpdateEvent", mock.AnythingOfType("events.Event")).
		Return(events.Event{}, mockError).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/", strings.NewReader(string(mockJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.Error(t, mockError, h.UpdateData(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_UpdateData_Bind_Error(t *testing.T) {

	mockError := errors.New("code=400, message=Unmarshal type error: expected=events.Event, got=array, offset=1")

	mockJSON := `[{"error_1":"01","error_2":"02"}]`

	mockResponse := module.Response{
		Status:  http.StatusBadRequest,
		Message: mockError.Error(),
		Data:    nil,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/", strings.NewReader(string(mockJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.Error(t, mockError, h.UpdateData(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_DeleteData(t *testing.T) {

	mockResponse := module.Response{
		Status:  http.StatusOK,
		Message: module.SUCCESS_DELETE_DATA,
		Data:    nil,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("DeleteEvent", "EVENT.01").
		Return(nil).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.DeleteData(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}

func TestEventHandler_DeleteData_Error(t *testing.T) {

	mockError := module.NOT_FOUND_ERROR

	mockResponse := module.Response{
		Status:  http.StatusNotFound,
		Message: mockError.Error(),
		Data:    nil,
	}

	mockResponseJSON, _ := json.Marshal(mockResponse)
	mockResponseString := string(mockResponseJSON)

	mockRepo := new(repository.MockEventRepository)
	mockRepo.On("DeleteEvent", "EVENT.01").
		Return(mockError).Once()

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events/:id")
	c.SetParamNames("id")
	c.SetParamValues("EVENT.01")

	h := &eventHandler{repo: mockRepo}

	// Assertions
	if assert.Error(t, mockError, h.DeleteData(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, mockResponseString, rec.Body.String())
	}

	mockRepo.AssertExpectations(t)
}
