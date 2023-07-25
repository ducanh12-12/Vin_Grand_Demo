package model

import "time"

type Report struct {
	Id        string    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"column:title" json:"title"`
	UserId    int       `gorm:"column:user_id" json:"user_id"`
	Content   string    `gorm:"column:content" json:"content"`
	Status    bool      `gorm:"column:status" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
