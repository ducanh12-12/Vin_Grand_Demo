package events

import (
	implFor "base-go/application/requires"
	"base-go/domain/model"
	"base-go/services/requires"
	"context"
)

type eventsServiceImpl struct {
	eventRepo requires.EventRepository
}

func NewEventService(eventRepo requires.EventRepository) implFor.EventsService {
	return &eventsServiceImpl{eventRepo}
}
func (svc *eventsServiceImpl) CreateEvent(ctx context.Context, event model.Event) (*model.Event, error) {
	return svc.eventRepo.CreateEvent(ctx, event)
}

func (svc *eventsServiceImpl) GetEvents(ctx context.Context) (*[]model.Event, error) {
	return svc.eventRepo.GetEvents(ctx)
}
func (svc *eventsServiceImpl) Retrieve(ctx context.Context, id int) (*model.Event, error) {
	return svc.eventRepo.Retrieve(ctx, id)
}
