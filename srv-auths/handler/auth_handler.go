package handler

import (
	"github.com/dfibrinogen/dfibrinogen-api/common"
	"github.com/dfibrinogen/dfibrinogen-api/srv-auths/model"
	"github.com/dfibrinogen/dfibrinogen-api/srv-auths/repository"
	"github.com/labstack/echo"
	"net/http"
)

type IAuthHandler interface {
	UserLogin(c echo.Context) error
	UserRegister(c echo.Context) error
}

type authHandler struct {
	repo repository.IAuthRepository
}

func NewAuthHandler(e *echo.Echo, repo repository.IAuthRepository) {

	handler := &authHandler{repo: repo}

	e.POST("/login", handler.UserLogin)
	e.POST("/register", handler.UserRegister)
}

func (h *authHandler) UserLogin(c echo.Context) error {

	var auth model.Auth

	err := c.Bind(&auth)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &common.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := h.repo.Login(auth)

	if err != nil {
		return c.JSON(http.StatusNotFound, &common.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &common.Response{
		Status:  http.StatusOK,
		Message: common.SUCCESS_LOGIN_DATA,
		Data:    result,
	})
}

func (h *authHandler) UserRegister(c echo.Context) error {

	var auth model.Auth

	err := c.Bind(&auth)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &common.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := h.repo.Register(auth)

	if err != nil {
		return c.JSON(http.StatusNotFound, &common.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &common.Response{
		Status:  http.StatusOK,
		Message: common.SUCCESS_REGISTER_DATA,
		Data:    result,
	})
}
