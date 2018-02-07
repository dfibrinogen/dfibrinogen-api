package model

import "time"

type DUser struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	Username  string     `json:"username" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	FullName  string     `json:"full_name" gorm:"-"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
