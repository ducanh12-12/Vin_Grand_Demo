package model

import "time"

type Notification struct {
	Id          int       `gorm:"column:id;primaryKey;autoIncrement" json:"id" validate:"required"`
	CreatetorId int       `gorm:"column:creator_id" json:"creator_id"`
	Title       string    `gorm:"column:title" json:"title"`
	Type        int       `gorm:"column:type" json:"type"`
	Data        string    `gorm:"column:data" json:"data"`
	Status      bool      `gorm:"column:status" json:"status"`
	UserId      int       `gorm:"column:userId" json:"userId"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
