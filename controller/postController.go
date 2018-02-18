package controller

import (
	"github.com/dafian47/dfibrinogen-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dafian47/dfibrinogen-api/util"
)

func (c *BaseController) GetPostAll(context *gin.Context) {

	var postList []model.DPost

	page := context.DefaultQuery("page", "1")
	perPage := context.DefaultQuery("per_page", "5")
	orderBy := context.DefaultQuery("order_by", "id ASC")

	limit, offset := util.GetLimitAndOffset(perPage, page)

	c.DB.Limit(limit).Offset(offset).Order(orderBy).Find(&postList)

	if len(postList) == 0 {
		responseJSON(context, http.StatusNotFound, "Not post at all", nil)
		return
	}

	for i := 0; i < len(postList); i++ {

		post := postList[i]

		var countLike int
		var countComment int

		c.DB.Model(&model.DPostLike{}).Where(&model.DPostLike{PostID:post.ID}).Count(&countLike)
		c.DB.Model(&model.DPostComment{}).Where(&model.DPostComment{PostID:post.ID}).Count(&countComment)

		postList[i].CountLike = countLike
		postList[i].CountComment = countComment
	}

	responseJSON(context, http.StatusOK, "Success get data forum post", postList)
}

func (c *BaseController) GetPostByID(context *gin.Context) {

	var post model.DPost

	postId := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DPost{ID: postId}).First(&post)

	if post.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data forum post", nil)
		return
	}

	var countLike int
	var countComment int

	c.DB.Model(&model.DPostLike{}).Where(&model.DPostLike{PostID:post.ID}).Count(&countLike)
	c.DB.Model(&model.DPostComment{}).Where(&model.DPostComment{PostID:post.ID}).Count(&countComment)

	post.CountLike = countLike
	post.CountComment = countComment

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

	responseJSON(context, http.StatusCreated, "Success save data forum post", post)
}

func (c *BaseController) UpdatePost(context *gin.Context) {

	var post model.DPost
	var postTemp model.DPost

	postId := util.ConvertStringToUint(context.Param("id"))

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

	postId := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DPost{ID: postId}).First(&post)

	if post.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data forum post", post)
		return
	}

	c.DB.Delete(&post)

	responseJSON(context, http.StatusOK, "Success delete data forum post", nil)
}

func (c *BaseController) AddViewer(context *gin.Context) {

	var post model.DPost
	var viewer int

	postID := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DPost{ID: postID}).First(&post)

	if post.ID == 0 {
		responseJSON(context, http.StatusOK, "Failed update count viewer post", nil)
		return
	}

	viewer = post.CountViewer
	viewer++

	c.DB.Model(&post).Where(&model.DPost{ID:postID}).Update("count_viewer", viewer)

	if post.CountViewer != viewer {
		responseJSON(context, http.StatusBadRequest, "Failed update count viewer post", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success update count viewer post", nil)
}

func (c *BaseController) AddLike(context *gin.Context) {

	var like model.DPostLike

	err := context.BindJSON(&like)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data like", nil)
		return
	}

	c.DB.Save(&like)

	if like.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed save data like", nil)
		return
	}

	responseJSON(context, http.StatusOK, "Success save data like", like)
}

func (c *BaseController) DeleteLike(context *gin.Context) {

	var like model.DPostLike

	err := context.BindJSON(&like)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data like", nil)
		return
	}

	c.DB.Where(&model.DPostLike{PostID:like.PostID, UserID:like.UserID}).Delete(&like)
	
	responseJSON(context, http.StatusOK, "Success delete data like", like)
}

func (c *BaseController) GetCommentByPostID(context *gin.Context) {
	
	var commentList []model.DPostComment

	postID := util.ConvertStringToUint(context.Param("id"))
	
	c.DB.Where(&model.DPostComment{PostID:postID}).Find(&commentList)

	if len(commentList) == 0 {
		responseJSON(context, http.StatusNotFound, "No data comment at all", nil)
		return 
	}
	
	responseJSON(context, http.StatusOK, "Success get data comment", commentList)
}

func (c *BaseController) AddComment(context *gin.Context) {

	var comment model.DPostComment
	
	err := context.BindJSON(&comment)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data comment", nil)
		return 
	}
	
	c.DB.Save(&comment)

	if comment.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed save data comment", nil)
		return 
	}
	
	responseJSON(context, http.StatusCreated, "Success save data comment", comment)
}

func (c *BaseController) UpdateComment(context *gin.Context) {

	var comment model.DPostComment
	var commentTemp model.DPostComment

	err := context.BindJSON(&comment)
	if err != nil {
		responseJSON(context, http.StatusBadRequest, "Failed bind data comment", nil)
		return
	}
	
	postID := util.ConvertStringToUint(context.Param("id"))
	
	c.DB.Where(&model.DPostComment{PostID:postID}).First(&commentTemp)

	if commentTemp.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data comment", nil)
		return 
	}
	
	c.DB.Where(&model.DPostComment{PostID:postID}).Save(&comment)

	if comment.ID == 0 {
		responseJSON(context, http.StatusBadRequest, "Failed update data comment", nil)
		return 
	}
	
	responseJSON(context, http.StatusOK, "Success update data comment", comment)
}

func (c *BaseController) DeleteComment(context *gin.Context) {

	var comment model.DPostComment

	commentID := util.ConvertStringToUint(context.Param("id"))

	c.DB.Where(&model.DPostComment{ID:commentID}).First(&comment)

	if comment.ID == 0 {
		responseJSON(context, http.StatusNotFound, "Not found data comment", nil)
		return
	}
	
	c.DB.Where(&model.DPostComment{ID:commentID}).Delete(&comment)
	
	responseJSON(context, http.StatusOK, "Success delete data comment", nil)
}