package repository

import (
	"github.com/dfibrinogen/dfibrinogen-api/events-service/model"
	"github.com/stretchr/testify/mock"
)

type EventRepo struct {
	mock.Mock
}

func InitMockEventRepo() *EventRepo {
	s := new(EventRepo)
	return s
}

func (m *EventRepo) FetchEventAll() ([]*model.Event, error) {
	ret := m.Called()

	var r0 []*model.Event
	if rf, ok := ret.Get(0).(func() []*model.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Event)
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

func (m *EventRepo) CreateEvent(event *model.Event) error {
	ret := m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Event) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
