package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/module"
	"github.com/dfibrinogen/dfibrinogen-api/module/auths"
	"github.com/jinzhu/gorm"
)

type IAuthRepository interface {
	Login(data auths.Auth) (auths.Auth, error)
	Register(data auths.Auth) (auths.Auth, error)
}

type authRepository struct {
	db *gorm.DB
}

func InitAuthRepo(db *gorm.DB) IAuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Login(data auths.Auth) (auths.Auth, error) {

	var auth auths.Auth

	r.db.Where(&auths.Auth{Username: data.Username}).First(&auth)

	if auth.AuthID == "" {
		return auths.Auth{}, module.NOT_FOUND_ERROR
	}

	return auth, nil
}

func (r *authRepository) Register(data auths.Auth) (auths.Auth, error) {

	var countUser int

	r.db.Where(&auths.Auth{Username: data.Username}).Count(&countUser)

	if countUser != 0 {
		return auths.Auth{}, module.CONFLIT_ERROR
	}

	r.db.Save(&data)

	if data.AuthID == "" {
		return auths.Auth{}, module.FAILED_SAVE_ERROR
	}

	return data, nil
}
