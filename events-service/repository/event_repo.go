package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/events-service/model"
)

type IEventRepository interface {
	GetDataAll() ([]model.Event, error)
	//GetDataByID(eventID string) (model.Event, error)
	//CreateData(data model.Event) (model.Event, error)
	//UpdateData(data model.Event) (model.Event, error)
	//DeleteData(eventID string) error
}

type eventRepository struct {

}

func InitEventRepo() IEventRepository {
	return &eventRepository{}
}

func (r *eventRepository) GetDataAll() ([]model.Event, error) {

}

