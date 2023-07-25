package presenter

import (
	"base-go/domain/model"
)

func Event(event *model.Event) model.Event {
	return model.Event{
		Title:       event.Title,
		Content:     event.Content,
		Description: event.Description,
		Avatar:      event.Avatar,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
}
