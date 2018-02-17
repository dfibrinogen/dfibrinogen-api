package controller

import (
	"github.com/dafian47/dfibrinogen-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"unicode"
)

func (c *BaseController) GetPostAll(context *gin.Context) {

	var postList []model.DPost

	c.DB.Find(&postList)

	if len(postList) == 0 {
		responseJSON(context, http.StatusNotFound, "Not post at all", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data forum post", postList)
}

func (c *BaseController) GetPostByID(context *gin.Context) {

	var post model.DPost

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	postId := uint(uid64)

	c.DB.Where(&model.DPost{ID: postId}).First(&post)

	if post.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data forum post", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success get data forum post", post)
}

func (c *BaseController) AddPost(context *gin.Context) {

	var post model.DPost

	err := context.BindJSON(&post)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data forum post", nil)
		return
	}

	c.DB.Save(&post)

	if post.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed save data forum post", post)
		return
	}

	responseJSON(context, http.StatusCreated, "Success save data forum post", unicode.Po)
}

func (c *BaseController) UpdatePost(context *gin.Context) {

	var post model.DPost
	var postTemp model.DPost

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	postId := uint(uid64)

	c.DB.Where(&model.DPost{ID: postId}).First(&postTemp)

	if postTemp.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data forum post", nil)
		return
	}

	err := context.BindJSON(&post)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data forum post", nil)
		return
	}

	c.DB.Where(&model.DPost{ID: postId}).Save(&post)

	if post.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed update data forum post", post)
		return
	}

	responseJSON(context, http.StatusOK, "Success update data forum post", post)
}

func (c *BaseController) DeletePost(context *gin.Context) {

	var post model.DPost

	id := context.Param("id")
	uid64, _ := strconv.ParseUint(id, 10, 32)
	postId := uint(uid64)

	c.DB.Where(&model.DPost{ID: postId}).First(&post)

	if post.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data forum post", post)
		return
	}

	c.DB.Delete(&post)

	responseJSON(context, http.StatusOK, "Success delete data forum post", nil)
}
