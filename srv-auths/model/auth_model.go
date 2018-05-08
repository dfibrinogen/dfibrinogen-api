package model

type Auth struct {
	AuthID   string `json:"auth_id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique_index"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
