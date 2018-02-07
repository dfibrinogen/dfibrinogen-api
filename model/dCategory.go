package model

import "time"

type DCategory struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Name      string     `json:"name" binding:"required"`
	UrlImage  string     `json:"url_image"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
