package controller

import (
	"github.com/dafian47/dfibrinogen-api/model"
	"github.com/dafian47/dfibrinogen-api/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *BaseController) UserLogin(context *gin.Context) {

	var user model.DUser
	var profile model.DProfile

	err := context.BindJSON(&user)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data user", nil)
		return
	}

	username := user.Username
	password := user.Password

	c.DB.Where(&model.DUser{Username: username}).First(&user)

	if user.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data user", nil)
		return
	}

	isMatched := util.MatchString(user.Password, password)
	if !isMatched {
		responseJSON(context, http.StatusBadRequest, "Wrong password", nil)
		return
	}

	c.DB.Where(&model.DProfile{UserID: user.ID}).First(&profile)

	if profile.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data profile", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data user", profile)
}

func (c *BaseController) UserRegister(context *gin.Context) {

	var user model.DUser
	var profile model.DProfile

	err := context.BindJSON(&user)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data user", nil)
		return
	}

	c.DB.Where(&model.DUser{Username: user.Username}).First(&profile)

	if user.ID != 0 {
		responseJSON(context, http.StatusBadRequest, "Username already exists", nil)
		return
	}

	fullName := user.FullName
	user.Password, _ = util.HashString(user.Password)

	c.DB.Save(&user)

	if user.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed save data user", nil)
		return
	}

	profile.UserID = user.ID
	profile.FullName = fullName

	c.DB.Save(&profile)

	if profile.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed save data profile", nil)
		return
	}

	responseJSON(context, http.StatusCreated, "Success save data user", profile)
}
