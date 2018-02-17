package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/dafian47/dfibrinogen-api/model"
)

type BaseController struct {
	DB *gorm.DB
}

func responseJSON(context *gin.Context, status int, message string, data interface{})  {
	context.JSON(status, &model.Response{
		Status:status,
		Message:message,
		Data:data,
	})
}
