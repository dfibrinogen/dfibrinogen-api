package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/common"
	"github.com/dfibrinogen/dfibrinogen-api/srv-events/model"
	"github.com/jinzhu/gorm"
)

type IEventRepository interface {
	FetchEventAll() ([]model.Event, error)
	FetchEventByID(id string) (model.Event, error)
	CreateEvent(data model.Event) (model.Event, error)
	UpdateEvent(data model.Event) (model.Event, error)
	DeleteEvent(id string) error
}

type eventRepository struct {
	db *gorm.DB
}

func InitEventRepo(db *gorm.DB) IEventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) FetchEventAll() ([]model.Event, error) {

	var dataList []model.Event

	r.db.Find(&dataList)

	if len(dataList) == 0 {
		return nil, common.EMPTY_ERROR
	}

	return dataList, nil
}

func (r *eventRepository) FetchEventByID(id string) (model.Event, error) {

	var data model.Event

	r.db.Where(&model.Event{ID: id}).
		First(&data)

	if data.ID == "" {
		return data, common.NOT_FOUND_ERROR
	}

	return data, nil
}

func (r *eventRepository) CreateEvent(data model.Event) (model.Event, error) {

	r.db.Save(&data)

	if data.ID == "" {
		return data, common.FAILED_SAVE_ERROR
	}

	return data, nil
}

func (r *eventRepository) UpdateEvent(data model.Event) (model.Event, error) {

	r.db.Save(&data)

	if data.ID == "" {
		return data, common.FAILED_UPDATE_ERROR
	}

	return data, nil
}

func (r *eventRepository) DeleteEvent(id string) error {

	var data model.Event

	r.db.Where(&model.Event{ID: id}).First(&data)

	if data.ID == "" {
		return common.NOT_FOUND_ERROR
	}

	r.db.Delete(&data)

	return nil
}
