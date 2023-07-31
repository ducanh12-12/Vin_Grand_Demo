package presenter

import (
	"base-go/domain/model"
)

func Event(event *model.Event) model.Event {
	return model.Event{
		Id:          event.Id,
		Title:       event.Title,
		Content:     event.Content,
		Description: event.Description,
		Avatar:      event.Avatar,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Status:      event.Status,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
}
