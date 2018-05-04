package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/events-service/model"
	"github.com/jinzhu/gorm"
)

type IEventRepository interface {
	FetchEventAll() ([]model.Event, error)
	FetchEventByID(id string) (model.Event, error)
	AddEvent(data model.Event) (model.Event, error)
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

	return dataList, nil
}

func (r *eventRepository) FetchEventByID(id string) (model.Event, error) {

	var data model.Event

	r.db.Where(&model.Event{ID: id}).
		First(&data)

	return data, nil
}

func (r *eventRepository) AddEvent(data model.Event) (model.Event, error) {

	r.db.Save(&data)

	return data, nil
}
