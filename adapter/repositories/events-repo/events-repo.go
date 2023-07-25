package events_repo

import (
	"base-go/domain/model"
	implFor "base-go/services/requires"
	"context"

	"gorm.io/gorm"
)

type eventsRepo struct {
	db *gorm.DB
}

func NewEventsRepo(db *gorm.DB) implFor.EventRepository {
	return &eventsRepo{db: db}
}

func (r *eventsRepo) CreateEvent(ctx context.Context, event model.Event) (*model.Event, error) {
	err := r.db.Create(&event).Error
	return &event, err
}

func (r *eventsRepo) Retrieve(ctx context.Context, id int) (*model.Event, error) {
	event := model.Event{}
	err := r.db.Where("id = ?", id).First(&event).Error
	return &event, err
}
func (r *eventsRepo) GetEvents(ctx context.Context) (*[]model.Event, error) {
	var event []model.Event
	err := r.db.Find(&event).Error
	return &event, err
}
