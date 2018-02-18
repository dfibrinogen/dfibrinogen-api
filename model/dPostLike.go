package model

type DPostLike struct {
	ID     uint `gorm:"primary_key" json:"id"`
	PostID uint `json:"post_id" binding:"required"`
	UserID uint `json:"user_id" binding:"required"`
}
