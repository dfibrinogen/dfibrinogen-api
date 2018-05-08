package handler

import (
	"testing"
	"github.com/labstack/echo"
	"github.com/dfibrinogen/dfibrinogen-api/srv-users/repository"
)

func TestNewUserHandler(t *testing.T) {

	mockRepo := new(repository.MockIUserRepository)

	e := echo.New()
	g := e.Group("/api/v1")
	NewUserHandler(g, mockRepo)
}

func TestUserHandler_GetDataAll(t *testing.T) {

}

func TestUserHandler_GetDataAll_Error(t *testing.T) {

}

func TestUserHandler_GetDataByID(t *testing.T) {

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
