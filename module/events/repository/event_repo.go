package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/module"
	"github.com/dfibrinogen/dfibrinogen-api/module/events"
	"github.com/jinzhu/gorm"
)

type IEventRepository interface {
	FetchEventAll() ([]events.Event, error)
	FetchEventByID(id string) (events.Event, error)
	CreateEvent(data events.Event) (events.Event, error)
	UpdateEvent(data events.Event) (events.Event, error)
	DeleteEvent(id string) error
}

type eventRepository struct {
	db *gorm.DB
}

func InitEventRepo(db *gorm.DB) IEventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) FetchEventAll() ([]events.Event, error) {

	var dataList []events.Event

	r.db.Find(&dataList)

	if len(dataList) == 0 {
		return nil, module.EMPTY_ERROR
	}

	return dataList, nil
}

func (r *eventRepository) FetchEventByID(id string) (events.Event, error) {

	var data events.Event

	r.db.Where(&events.Event{ID: id}).
		First(&data)

	if data.ID == "" {
		return data, module.NOT_FOUND_ERROR
	}

	return data, nil
}

func (r *eventRepository) CreateEvent(data events.Event) (events.Event, error) {

	r.db.Save(&data)

	if data.ID == "" {
		return data, module.FAILED_SAVE_ERROR
	}

	return data, nil
}

func (r *eventRepository) UpdateEvent(data events.Event) (events.Event, error) {

	r.db.Save(&data)

	if data.ID == "" {
		return data, module.FAILED_UPDATE_ERROR
	}

	return data, nil
}

func (r *eventRepository) DeleteEvent(id string) error {

	var data events.Event

	r.db.Where(&events.Event{ID: id}).First(&data)

	if data.ID == "" {
		return module.NOT_FOUND_ERROR
	}

	r.db.Delete(&data)

	return nil
}
