package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/srv-events/model"
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

func (m *EventRepo) CreateEvent(data model.Event) (model.Event, error) {
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

func (m *EventRepo) UpdateEvent(data model.Event) (model.Event, error) {
	ret := m.Called(data)

	var r0 model.Event
	if rf, ok := ret.Get(0).(func(model.Event) model.Event); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(model.Event)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Event) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *EventRepo) DeleteEvent(id string) error {
	ret := m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
