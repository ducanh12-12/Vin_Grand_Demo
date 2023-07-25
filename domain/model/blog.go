package model

import "time"

type Blog struct {
	Id          string    `gorm:"column:id;primaryKey" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Content     string    `gorm:"column:content" json:"content"`
	Avatar      string    `gorm:"column:avatar" json:"avatar"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
