package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/common"
	"github.com/dfibrinogen/dfibrinogen-api/srv-auths/model"
	"github.com/jinzhu/gorm"
)

type IAuthRepository interface {
	Login(data model.Auth) (model.Auth, error)
	Register(data model.Auth) (model.Auth, error)
}

type authRepository struct {
	db *gorm.DB
}

func InitAuthRepo(db *gorm.DB) IAuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Login(data model.Auth) (model.Auth, error) {

	var auth model.Auth

	r.db.Where(&model.Auth{Username: data.Username}).First(&auth)

	if auth.AuthID == "" {
		return model.Auth{}, common.NOT_FOUND_ERROR
	}

	return auth, nil
}

func (r *authRepository) Register(data model.Auth) (model.Auth, error) {

	var countUser int

	r.db.Where(&model.Auth{Username: data.Username}).Count(&countUser)

	if countUser != 0 {
		return model.Auth{}, common.CONFLIT_ERROR
	}

	r.db.Save(&data)

	if data.AuthID == "" {
		return model.Auth{}, common.FAILED_SAVE_ERROR
	}

	return data, nil
}
