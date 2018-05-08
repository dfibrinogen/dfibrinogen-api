package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/common"
	"github.com/dfibrinogen/dfibrinogen-api/srv-users/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	FetchUserAll() ([]model.User, error)
	FetchUserByID(id string) (model.User, error)
	CreateUser(data model.User) (model.User, error)
	UpdateUser(data model.User) (model.User, error)
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func InitUserRepo(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FetchUserAll() ([]model.User, error) {

	var dataList []model.User

	r.db.Find(&dataList)

	if len(dataList) == 0 {
		return nil, common.EMPTY_ERROR
	}

	return dataList, nil
}

func (r *userRepository) FetchUserByID(id string) (model.User, error) {

	var data model.User

	r.db.Where(&model.User{ID: id}).
		First(&data)

	if data.ID == "" {
		return data, common.NOT_FOUND_ERROR
	}

	return data, nil
}

func (r *userRepository) CreateUser(data model.User) (model.User, error) {

	r.db.Save(&data)

	if data.ID == "" {
		return data, common.FAILED_SAVE_ERROR
	}

	return data, nil
}

func (r *userRepository) UpdateUser(data model.User) (model.User, error) {

	r.db.Save(&data)

	if data.ID == "" {
		return data, common.FAILED_UPDATE_ERROR
	}

	return data, nil
}

func (r *userRepository) DeleteUser(id string) error {

	var data model.User

	r.db.Where(&model.User{ID: id}).First(&data)

	if data.ID == "" {
		return common.NOT_FOUND_ERROR
	}

	r.db.Delete(&data)

	return nil
}
