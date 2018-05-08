package model

import "time"

type User struct {
	ID        string     `json:"id" gorm:"primary_key"`
	Name      string     `json:"name"`
	Gender    string     `json:"gender"`
	Age       int        `json:"age"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
