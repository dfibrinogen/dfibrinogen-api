package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/module"
	"github.com/dfibrinogen/dfibrinogen-api/module/users"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	FetchUserAll() ([]users.User, error)
	FetchUserByID(id string) (users.User, error)
	CreateUser(data users.User) (users.User, error)
	UpdateUser(data users.User) (users.User, error)
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func InitUserRepo(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FetchUserAll() ([]users.User, error) {

	var dataList []users.User

	r.db.Find(&dataList)

	if len(dataList) == 0 {
		return nil, module.EMPTY_ERROR
	}

	return dataList, nil
}

func (r *userRepository) FetchUserByID(id string) (users.User, error) {

	var data users.User

	r.db.Where(&users.User{ID: id}).
		First(&data)

	if data.ID == "" {
		return data, module.NOT_FOUND_ERROR
	}

	return data, nil
}

func (r *userRepository) CreateUser(data users.User) (users.User, error) {

	r.db.Save(&data)

	if data.ID == "" {
		return data, module.FAILED_SAVE_ERROR
	}

	return data, nil
}

func (r *userRepository) UpdateUser(data users.User) (users.User, error) {

	r.db.Save(&data)

	if data.ID == "" {
		return data, module.FAILED_UPDATE_ERROR
	}

	return data, nil
}

func (r *userRepository) DeleteUser(id string) error {

	var data users.User

	r.db.Where(&users.User{ID: id}).First(&data)

	if data.ID == "" {
		return module.NOT_FOUND_ERROR
	}

	r.db.Delete(&data)

	return nil
}
