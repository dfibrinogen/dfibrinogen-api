package controller

import (
	"github.com/dafian47/dfibrinogen-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *BaseController) GetUserAll(context *gin.Context) {

	var profileList []model.DProfile

	c.DB.Find(&profileList)

	if len(profileList) == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not user at all",
		})
		return
	}

	context.JSON(http.StatusOK, profileList)
}

func (c *BaseController) GetProfileByID(context *gin.Context) {

	var profile model.DProfile

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	userId := uint(uid64)

	c.DB.Where(&model.DProfile{UserID: userId}).First(&profile)

	if profile.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not found user",
		})
		return
	}

	context.JSON(http.StatusOK, profile)
}

func (c *BaseController) UpdateProfile(context *gin.Context) {

	var profile model.DProfile
	var profileTemp model.DProfile

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	userId := uint(uid64)

	c.DB.Where(&model.DProfile{UserID: userId}).First(&profileTemp)

	if profileTemp.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not found user",
		})
		return
	}

	err := context.BindJSON(&profile)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed bind data profile",
		})
		return
	}

	c.DB.Where(&model.DProfile{UserID: userId}).Save(&profile)

	if profile.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed update profile",
			"data":    profile,
		})
		return
	}

	context.JSON(http.StatusOK, profile)
}

func (c *BaseController) DeleteUser(context *gin.Context) {

	var user model.DUser
	var profile model.DProfile

	// TODO Delete Data Post with User Id

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	userId := uint(uid64)

	c.DB.Where(&model.DUser{ID: userId}).First(&user)
	c.DB.Where(&model.DProfile{UserID: userId}).First(&profile)

	if user.ID == 0 || profile.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not found data",
		})
		return
	}

	c.DB.Delete(&user)
	c.DB.Delete(&profile)

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success delete user profile",
	})
}
