package model

import "time"

type DProfile struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	FullName  string     `json:"full_name" binding:"required"`
	Gender    string     `json:"gender"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
	UrlImage  string     `json:"url_image"`
	UserID    uint       `json:"user_id" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
