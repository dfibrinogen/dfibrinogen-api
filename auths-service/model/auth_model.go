package model

type Auth struct {
	AuthID   string `json:"auth_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
