package controller

import (
	"github.com/dafian47/dfibrinogen-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *BaseController) GetPostAll(context *gin.Context) {

	var postList []model.DPost

	c.DB.Find(&postList)

	if len(postList) == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not post at all",
		})
		return
	}

	context.JSON(http.StatusOK, postList)
}

func (c *BaseController) GetPostByID(context *gin.Context) {

	var post model.DPost

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	postId := uint(uid64)

	c.DB.Where(&model.DPost{ID: postId}).First(&post)

	if post.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No found post",
		})
		return
	}

	context.JSON(http.StatusOK, post)
}

func (c *BaseController) AddPost(context *gin.Context) {

	var post model.DPost

	err := context.BindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed bind data post",
		})
		return
	}

	c.DB.Save(&post)

	if post.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed save post",
		})
		return
	}

	context.JSON(http.StatusCreated, post)
}

func (c *BaseController) UpdatePost(context *gin.Context) {

	var post model.DPost
	var postTemp model.DPost

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	postId := uint(uid64)

	c.DB.Where(&model.DPost{ID: postId}).First(&postTemp)

	if postTemp.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No found post",
		})
		return
	}

	err := context.BindJSON(&post)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed bind data post",
		})
		return
	}

	c.DB.Where(&model.DPost{ID: postId}).Save(&post)

	if post.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Failed update post",
			"data":    post,
		})
		return
	}

	context.JSON(http.StatusOK, post)
}

func (c *BaseController) DeletePost(context *gin.Context) {

	var post model.DPost

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	postId := uint(uid64)

	c.DB.Where(&model.DPost{ID: postId}).First(&post)

	if post.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No found post",
		})
		return
	}

	c.DB.Delete(&post)

	context.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success delete post",
	})
}
