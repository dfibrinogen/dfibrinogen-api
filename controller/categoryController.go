package controller

import (
	"github.com/dafian47/dfibrinogen-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dafian47/dfibrinogen-api/util"
)

func (c *BaseController) GetCategoryAll(context *gin.Context) {

	var categoryList []model.DCategory

	c.DB.Find(&categoryList)

	if len(categoryList) == 0 {
		responseJSON(context, http.StatusNotFound, "Not category at all", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data forum category", categoryList)
}

func (c *BaseController) GetCategoryByID(context *gin.Context) {

	var category model.DCategory

	categoryId := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DCategory{ID: categoryId}).First(&category)

	if category.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data forum category", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data forum category", category)
}

func (c *BaseController) AddCategory(context *gin.Context) {

	var category model.DCategory

	err := context.BindJSON(&category)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data forum category", nil)
		return
	}

	c.DB.Save(&category)

	if category.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed save data forum category", nil)
		return
	}

	responseJSON(context, http.StatusCreated, "Success save data forum category", category)
}

func (c *BaseController) UpdateCategory(context *gin.Context) {

	var category model.DCategory
	var categoryTemp model.DCategory

	categoryId := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DCategory{ID: categoryId}).First(&categoryTemp)

	if categoryTemp.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data forum category", nil)
		return
	}

	err := context.BindJSON(&category)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data forum category", nil)
		return
	}

	c.DB.Where(&model.DCategory{ID: categoryId}).Save(&category)

	if categoryTemp.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed update data forum category", category)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data forum category", category)
}

func (c *BaseController) DeleteCategory(context *gin.Context) {

	var category model.DCategory

	// TODO Delete Data Post with Category Id

	categoryId := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DCategory{ID: categoryId}).First(&category)

	if category.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data forum category", nil)
		return
	}

	c.DB.Delete(&category)

	responseJSON(context, http.StatusOK, "Success delete data forum category", nil)
}
