package service

import (
	"github.com/dfibrinogen/dfibrinogen-api/events-service/repository"
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

func NewEventService(e *echo.Group, repo repository.IEventRepository)  {

	handler := &eventService{repo:repo}

	e.GET("/events", handler.GetDataAll)
}

func (s *eventService) GetDataAll(c echo.Context) error {

	dataList, err := s.repo.GetDataAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, dataList)
}