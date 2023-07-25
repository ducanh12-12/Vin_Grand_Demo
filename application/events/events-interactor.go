package events

import (
	"base-go/application/requires"
	"base-go/common/logger"
	"base-go/domain/model"
	"context"
	"time"
)

type EventsInteractor struct {
	eventService requires.EventsService
}

func NewEventsInteractor(eventService requires.EventsService) EventsInteractor {
	return EventsInteractor{eventService}
}
func (interactor *EventsInteractor) CreateEvent(ctx context.Context, event AddEventIpt) (*model.Event, error) {
	now := time.Now()
	newEvent := model.Event{
		Title:       event.Title,
		Content:     event.Content,
		Description: event.Description,
		Avatar:      event.Avatar,
		CreatedAt:   now,
	}
	eventNew, err := interactor.eventService.CreateEvent(ctx, newEvent)
	if err != nil {
		logger.Error("Unable to add event, error: %s", err.Error())
		return nil, err
	}
	return eventNew, nil
}
func (interactor *EventsInteractor) GetEvent(ctx context.Context, id int) (*GetEventResp, error) {
	return interactor.eventService.Retrieve(ctx, id)
}
func (interactor *EventsInteractor) GetEvents(ctx context.Context) (*[]GetEventResp, error) {
	return interactor.eventService.GetEvents(ctx)
}
