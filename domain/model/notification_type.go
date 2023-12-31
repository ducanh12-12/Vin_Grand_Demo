package model

import "time"

type NotificationType struct {
	Id        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id" validate:"required"`
	Title     string    `gorm:"column:title" json:"title"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
