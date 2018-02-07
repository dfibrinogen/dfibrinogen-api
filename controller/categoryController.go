package controller

import (
	"github.com/dafian47/dfibrinogen-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *BaseController) GetCategoryAll(context *gin.Context) {

	var categoryList []model.DCategory

	c.DB.Find(&categoryList)

	if len(categoryList) == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not category at all",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get category all",
		"data":    categoryList,
	})
}

func (c *BaseController) GetCategoryByID(context *gin.Context) {

	var category model.DCategory

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	categoryId := uint(uid64)

	c.DB.Where(&model.DCategory{ID: categoryId}).First(&category)

	if category.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not found category",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success get category",
		"data":    category,
	})
}

func (c *BaseController) AddCategory(context *gin.Context) {

	var category model.DCategory

	err := context.BindJSON(&category)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed bind data category",
		})
		return
	}

	c.DB.Save(&category)

	if category.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed save data category",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success save category",
		"data":    category,
	})
}

func (c *BaseController) UpdateCategory(context *gin.Context) {

	var category model.DCategory
	var categoryTemp model.DCategory

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	categoryId := uint(uid64)

	c.DB.Where(&model.DCategory{ID: categoryId}).First(&categoryTemp)

	if categoryTemp.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not found category",
		})
		return
	}

	err := context.BindJSON(&category)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed bind data category",
		})
		return
	}

	c.DB.Where(&model.DCategory{ID: categoryId}).Save(&category)

	if categoryTemp.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed update data category",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success update data category",
		"data":    category,
	})
}

func (c *BaseController) DeleteCategory(context *gin.Context) {

	var category model.DCategory

	// TODO Delete Data Post with Category Id

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	categoryId := uint(uid64)

	c.DB.Where(&model.DCategory{ID: categoryId}).First(&category)

	if category.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not found category",
		})
		return
	}

	c.DB.Delete(&category)

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success delete data category",
	})
}
