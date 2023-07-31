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
	err := r.db.Order("id asc").Find(&event).Error
	return &event, err
}
func (r *eventsRepo) Update(ctx context.Context, event model.Event, id int) (*model.Event, error) {
	eventOld := model.Event{}
	if err := r.db.Where("id = ?", id).First(&eventOld).Error; err != nil {
		return &eventOld, err
	}
	eventCopy := event
	err := r.db.Model(&eventOld).Updates(&eventCopy).Error
	return &eventOld, err
}
func (r *eventsRepo) Delete(ctx context.Context, id int) (string, error) {
	err := r.db.Delete(model.Event{}, id).Error
	return "okla", err
}
