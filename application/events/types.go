package events

import (
	"base-go/domain/model"
)

// command
type AddEventIpt struct {
	Title       string `json:"title"  validate:"required"`
	Content     string `json:"content"  validate:"required"`
	Description string `json:"description"`
	Avatar      string `json:"avatar" `
}

// query

// response
type GetEventResp = model.Event
