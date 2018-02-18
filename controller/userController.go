package controller

import (
	"github.com/dafian47/dfibrinogen-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dafian47/dfibrinogen-api/util"
)

func (c *BaseController) GetUserAll(context *gin.Context) {

	var profileList []model.DProfile

	c.DB.Find(&profileList)

	if len(profileList) == 0 {
		responseJSON(context, http.StatusNotFound, "Not user at all", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data user", profileList)
}

func (c *BaseController) GetProfileByID(context *gin.Context) {

	var profile model.DProfile

	userId := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DProfile{UserID: userId}).First(&profile)

	if profile.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data user", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data user", profile)
}

func (c *BaseController) UpdateProfile(context *gin.Context) {

	var profile model.DProfile
	var profileTemp model.DProfile

	userId := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DProfile{UserID: userId}).First(&profileTemp)

	if profileTemp.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data user", nil)
		return
	}

	err := context.BindJSON(&profile)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data user", nil)
		return
	}

	c.DB.Where(&model.DProfile{UserID: userId}).Save(&profile)

	if profile.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed update data user", profile)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data user", profile)
}

func (c *BaseController) DeleteUser(context *gin.Context) {

	var user model.DUser
	var profile model.DProfile

	// TODO Delete Data Post with User Id

	userId := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DUser{ID: userId}).First(&user)
	c.DB.Where(&model.DProfile{UserID: userId}).First(&profile)

	if user.ID == 0 || profile.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data user", nil)
		return
	}

	c.DB.Delete(&user)
	c.DB.Delete(&profile)

	responseJSON(context, http.StatusOK, "Success delete data user", nil)
}
