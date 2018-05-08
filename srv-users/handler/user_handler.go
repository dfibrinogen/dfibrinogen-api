package handler

import (
	"github.com/dfibrinogen/dfibrinogen-api/common"
	"github.com/dfibrinogen/dfibrinogen-api/srv-users/model"
	"github.com/dfibrinogen/dfibrinogen-api/srv-users/repository"
	"github.com/labstack/echo"
	"net/http"
)

type IUserHandler interface {
	GetDataAll(c echo.Context) error
	GetDataByID(c echo.Context) error
	CreateData(c echo.Context) error
	UpdateData(c echo.Context) error
	DeleteData(c echo.Context) error
}

type userHandler struct {
	repo repository.IUserRepository
}

func NewUserHandler(e *echo.Group, repo repository.IUserRepository) {

	handler := &userHandler{repo: repo}

	e.GET("/users", handler.GetDataAll)
	e.GET("/users/:id", handler.GetDataByID)
	e.POST("/users", handler.CreateData)
	e.PUT("/users/:id", handler.UpdateData)
	e.DELETE("/users/:id", handler.DeleteData)
}

func (h *userHandler) GetDataAll(c echo.Context) error {

	dataResults, err := h.repo.FetchUserAll()

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

func (h *userHandler) GetDataByID(c echo.Context) error {

	id := c.Param("id")

	dataResult, err := h.repo.FetchUserByID(id)

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

func (h *userHandler) CreateData(c echo.Context) error {

	var data model.User

	err := c.Bind(&data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &common.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	dataResult, err := h.repo.CreateUser(data)

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

func (h *userHandler) UpdateData(c echo.Context) error {

	var data model.User

	err := c.Bind(&data)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &common.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	dataResult, err := h.repo.UpdateUser(data)

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

func (h *userHandler) DeleteData(c echo.Context) error {

	id := c.Param("id")

	err := h.repo.DeleteUser(id)

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
