package events

import (
	"base-go/domain/model"
	"time"
)

// command
type AddEventIpt struct {
	Title       string    `json:"title"  validate:"required"`
	Content     string    `json:"content"  validate:"required"`
	Description string    `json:"description"`
	Avatar      string    `json:"avatar" `
	Status      bool      `gorm:"column:status" json:"status"`
	StartTime   time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime     time.Time `gorm:"column:end_time" json:"end_time"`
}

// query

// response
type GetEventResp = model.Event
