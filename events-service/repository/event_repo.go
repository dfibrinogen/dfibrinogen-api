package repository

import (
	"database/sql"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/model"
)

type IEventRepository interface {
	FetchEventAll() ([]*model.Event, error)
	CreateEvent(event *model.Event) error
}

type eventRepository struct {
	db *sql.DB
}

func InitEventRepo(db *sql.DB) IEventRepository {
	return &eventRepository{db: db}
}

func (r *eventRepository) FetchEventAll() ([]*model.Event, error) {

	var dataList []*model.Event

	rows, err := r.db.Query(
		"SELECT id, name, location, created_at FROM events WHERE deleted_at IS NULL ORDER BY created_at ASC ;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {

		data := &model.Event{}

		if err := rows.Scan(&data.ID, &data.Name, &data.Location, &data.CreatedAt); err != nil {
			return nil, err
		}

		dataList = append(dataList, data)
	}

	return dataList, nil
}

func (r *eventRepository) CreateEvent(event *model.Event) error {

	_, err := r.db.Query(
		"INSERT INTO events (id, name, location, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)",
		event.ID, event.Name, event.Location, event.CreatedAt, event.UpdatedAt, event.DeletedAt)
	return err
}
