package model

import "time"

type Event struct {
	Id          int       `gorm:"column:id;primaryKey;autoIncrement" json:"id" validate:"required"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Content     string    `gorm:"column:content" json:"content"`
	Avatar      string    `gorm:"column:avatar" json:"avatar"`
	Status      bool      `gorm:"column:status" json:"status"`
	StartTime   time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime     time.Time `gorm:"column:end_time" json:"end_time"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
