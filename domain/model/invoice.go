package model

import "time"

type Invoice struct {
	Id          string    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Address     string    `gorm:"column:avatar" json:"avatar"`
	Price       float64   `gorm:"column:name" json:"price"`
	Point       int       `gorm:"column:description" json:"point"`
	Description string    `gorm:"column:description" json:"description"`
	Status      bool      `gorm:"column:status" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
