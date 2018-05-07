package service

import (
	"github.com/dfibrinogen/dfibrinogen-api/events-service/model"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/repository"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/util"
	"github.com/labstack/echo"
	"net/http"
)

type IEventService interface {
	GetDataAll(c echo.Context) error
	GetDataByID(c echo.Context) error
	CreateData(c echo.Context) error
	UpdateData(c echo.Context) error
	DeleteData(c echo.Context) error
}

type eventService struct {
	repo repository.IEventRepository
}

func NewEventService(e *echo.Group, repo repository.IEventRepository) {

	handler := &eventService{repo: repo}

	e.GET("/events", handler.GetDataAll)
	e.GET("/events/:id", handler.GetDataByID)
	e.POST("/events", handler.CreateData)
}

func (s *eventService) GetDataAll(c echo.Context) error {

	dataResults, err := s.repo.FetchEventAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, &util.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &util.Response{
		Status:  http.StatusOK,
		Message: util.SUCCESS_GET_DATA_ALL,
		Data:    dataResults,
	})
}

func (s *eventService) GetDataByID(c echo.Context) error {

	id := c.Param("id")

	dataResult, err := s.repo.FetchEventByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &util.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &util.Response{
		Status:  http.StatusOK,
		Message: util.SUCCESS_GET_DATA,
		Data:    dataResult,
	})
}

func (s *eventService) CreateData(c echo.Context) error {

	var data model.Event

	err := c.Bind(&data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &util.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	dataResult, err := s.repo.CreateEvent(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &util.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, &util.Response{
		Status:  http.StatusCreated,
		Message: util.SUCCESS_CREATE_DATA,
		Data:    dataResult,
	})
}

func (s *eventService) UpdateData(c echo.Context) error {

	var data model.Event

	err := c.Bind(&data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &util.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	dataResult, err := s.repo.UpdateEvent(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &util.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &util.Response{
		Status:  http.StatusOK,
		Message: util.SUCCESS_UPDATE_DATA,
		Data:    dataResult,
	})
}

func (s *eventService) DeleteData(c echo.Context) error {

	id := c.Param("id")

	err := s.repo.DeleteEvent(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &util.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &util.Response{
		Status:  http.StatusOK,
		Message: util.SUCCESS_DELETE_DATA,
		Data:    nil,
	})
}
