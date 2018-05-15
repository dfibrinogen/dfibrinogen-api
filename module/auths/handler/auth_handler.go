package handler

import (
	"github.com/dfibrinogen/dfibrinogen-api/module"
	"github.com/dfibrinogen/dfibrinogen-api/module/auths"
	"github.com/dfibrinogen/dfibrinogen-api/module/auths/repository"
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

	var auth auths.Auth

	err := c.Bind(&auth)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &module.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := h.repo.Login(auth)

	if err != nil {
		return c.JSON(http.StatusNotFound, &module.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &module.Response{
		Status:  http.StatusOK,
		Message: module.SUCCESS_LOGIN_DATA,
		Data:    result,
	})
}

func (h *authHandler) UserRegister(c echo.Context) error {

	var auth auths.Auth

	err := c.Bind(&auth)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &module.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	result, err := h.repo.Register(auth)

	if err != nil {
		return c.JSON(http.StatusNotFound, &module.Response{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, &module.Response{
		Status:  http.StatusOK,
		Message: module.SUCCESS_REGISTER_DATA,
		Data:    result,
	})
}
