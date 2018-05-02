package service

import (
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
}

func (s *eventService) GetDataAll(c echo.Context) error {

	dataList, err := s.repo.FetchEventAll()

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
		Data:    dataList,
	})
}
