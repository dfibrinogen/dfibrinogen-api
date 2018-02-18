package model

import "time"

type DPostComment struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	PostID    uint       `json:"post_id" binding:"required"`
	UserID    uint       `json:"user_id" binding:"required"`
	Comment   string     `json:"comment"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
