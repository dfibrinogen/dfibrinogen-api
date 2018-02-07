package controller

import "github.com/jinzhu/gorm"

type BaseController struct {
	DB *gorm.DB
}
