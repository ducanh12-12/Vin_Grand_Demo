package model

import "time"

type Map struct {
	Id        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id" validate:"required"`
	Latitude  float64   `gorm:"column:latitude" json:"latitude"`
	Longitude float64   `gorm:"column:longitude" json:"longitude"`
	Address   string    `gorm:"column:address" json:"address"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
