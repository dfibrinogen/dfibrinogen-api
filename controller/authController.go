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
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed bind data login",
		})
		return
	}

	username := user.Username
	password := user.Password

	c.DB.Where(&model.DUser{Username: username}).First(&user)

	if user.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not user found",
		})
		return
	}

	isMatched := util.MatchString(user.Password, password)
	if !isMatched {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Wrong Password",
		})
		return
	}

	c.DB.Where(&model.DProfile{UserID: user.ID}).First(&profile)

	if profile.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not profile found",
		})
		return
	}

	context.JSON(http.StatusOK, profile)
}

func (c *BaseController) UserRegister(context *gin.Context) {

	var user model.DUser
	var profile model.DProfile

	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed bind data user",
		})
		return
	}

	c.DB.Where(&model.DUser{Username: user.Username}).First(&profile)

	if user.ID != 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Username already exists",
		})
		return
	}

	fullName := user.FullName
	user.Password, _ = util.HashString(user.Password)

	c.DB.Save(&user)

	if user.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed save data user",
		})
		return
	}

	profile.UserID = user.ID
	profile.FullName = fullName

	c.DB.Save(&profile)

	if profile.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed save data profile",
		})
		return
	}

	context.JSON(http.StatusCreated, profile)
}
