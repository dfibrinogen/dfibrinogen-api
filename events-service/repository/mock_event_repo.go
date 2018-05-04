package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/events-service/model"
	"github.com/stretchr/testify/mock"
)

type EventRepo struct {
	mock.Mock
}

func (m *EventRepo) FetchEventAll() ([]model.Event, error) {
	ret := m.Called()

	var r0 []model.Event
	if rf, ok := ret.Get(0).(func() []model.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *EventRepo) FetchEventByID(id string) (model.Event, error) {
	ret := m.Called()

	var r0 model.Event
	if rf, ok := ret.Get(0).(func() model.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *EventRepo) AddEvent(data model.Event) (model.Event, error) {
	ret := m.Called(data)

	var r0 model.Event
	if rf, ok := ret.Get(0).(func() model.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
