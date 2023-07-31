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
	layout := "2/1/2006"
	var startTime *time.Time
	var endTime *time.Time
	var err error
	parsedStartTime, err := time.Parse(layout, event.StartTime)
	if err == nil {
		startTime = &parsedStartTime
	}
	parsedEndTime, err := time.Parse(layout, event.EndTime)
	if err == nil {
		endTime = &parsedEndTime
	}
	newEvent := model.Event{
		Title:       event.Title,
		Content:     event.Content,
		Description: event.Description,
		Avatar:      event.Avatar,
		StartTime:   startTime,
		EndTime:     endTime,
		Status:      event.Status,
		CreatedAt:   now,
	}
	eventNew, err := interactor.eventService.CreateEvent(ctx, newEvent)
	if err != nil {
		logger.Error("Unable to add event, error: %s", err.Error())
		return nil, err
	}
	return eventNew, nil
}
func (interactor *EventsInteractor) Update(ctx context.Context, event UpdateEventIpt, id int) (*model.Event, error) {
	now := time.Now()
	layout := "2/1/2006"
	var startTime *time.Time
	var endTime *time.Time
	var err error
	parsedStartTime, err := time.Parse(layout, event.StartTime)
	if err == nil {
		startTime = &parsedStartTime
	}
	parsedEndTime, err := time.Parse(layout, event.EndTime)
	if err == nil {
		endTime = &parsedEndTime
	}
	newEvent := model.Event{
		Title:       event.Title,
		Content:     event.Content,
		Description: event.Description,
		Avatar:      event.Avatar,
		StartTime:   startTime,
		EndTime:     endTime,
		Status:      event.Status,
		CreatedAt:   now,
	}
	eventNew, err := interactor.eventService.Update(ctx, newEvent, id)
	if err != nil {
		logger.Error("Unable to add event, error: %s", err.Error())
		return nil, err
	}
	return eventNew, nil
}
func (interactor *EventsInteractor) GetEvent(ctx context.Context, id int) (*GetEventResp, error) {
	return interactor.eventService.Retrieve(ctx, id)
}
func (interactor *EventsInteractor) Delete(ctx context.Context, id int) (string, error) {
	return interactor.eventService.Delete(ctx, id)
}
func (interactor *EventsInteractor) GetEvents(ctx context.Context) (*[]GetEventResp, error) {
	return interactor.eventService.GetEvents(ctx)
}
