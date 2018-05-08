package handler

import (
	"github.com/dfibrinogen/dfibrinogen-api/common"
	"github.com/dfibrinogen/dfibrinogen-api/srv-events/model"
	"github.com/dfibrinogen/dfibrinogen-api/srv-events/repository"
	"github.com/labstack/echo"
	"net/http"
)

type IEventHandler interface {
	GetDataAll(c echo.Context) error
	GetDataByID(c echo.Context) error
	CreateData(c echo.Context) error
	UpdateData(c echo.Context) error
	DeleteData(c echo.Context) error
}

type eventHandler struct {
	repo repository.IEventRepository
}

func NewEventHandler(e *echo.Group, repo repository.IEventRepository) {

	handler := &eventHandler{repo: repo}

	e.GET("/events", handler.GetDataAll)
	e.GET("/events/:id", handler.GetDataByID)
	e.POST("/events", handler.CreateData)
	e.PUT("/events/:id", handler.UpdateData)
	e.DELETE("/events/:id", handler.DeleteData)
}

func (h *eventHandler) GetDataAll(c echo.Context) error {

	dataResults, err := h.repo.FetchEventAll()

	if err != nil {
		return c.JSON(http.StatusNotFound, &common.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &common.Response{
		Status:  http.StatusOK,
		Message: common.SUCCESS_GET_DATA_ALL,
		Data:    dataResults,
	})
}

func (h *eventHandler) GetDataByID(c echo.Context) error {

	id := c.Param("id")

	dataResult, err := h.repo.FetchEventByID(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, &common.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &common.Response{
		Status:  http.StatusOK,
		Message: common.SUCCESS_GET_DATA,
		Data:    dataResult,
	})
}

func (h *eventHandler) CreateData(c echo.Context) error {

	var data model.Event

	err := c.Bind(&data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &common.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	dataResult, err := h.repo.CreateEvent(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &common.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusCreated, &common.Response{
		Status:  http.StatusCreated,
		Message: common.SUCCESS_CREATE_DATA,
		Data:    dataResult,
	})
}

func (h *eventHandler) UpdateData(c echo.Context) error {

	var data model.Event

	err := c.Bind(&data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &common.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	dataResult, err := h.repo.UpdateEvent(data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &common.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &common.Response{
		Status:  http.StatusOK,
		Message: common.SUCCESS_UPDATE_DATA,
		Data:    dataResult,
	})
}

func (h *eventHandler) DeleteData(c echo.Context) error {

	id := c.Param("id")

	err := h.repo.DeleteEvent(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, &common.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &common.Response{
		Status:  http.StatusOK,
		Message: common.SUCCESS_DELETE_DATA,
		Data:    nil,
	})
}
