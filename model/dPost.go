package model

import "time"

type DPost struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	Title      string     `json:"title" binding:"required"`
	Content    string     `json:"content" binding:"required" sql:"type=text;"`
	UrlImage   string     `json:"url_image"`
	CategoryID uint       `json:"category_id" binding:"required"`
	UserID     uint       `json:"user_id" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
