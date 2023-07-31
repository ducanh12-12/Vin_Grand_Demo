package model

import "time"

type Invoice struct {
	Id          int       `gorm:"column:id;primaryKey;autoIncrement" json:"id" validate:"required"`
	Address     *string   `gorm:"column:address;not null" json:"address" validate:"required"`
	Price       float64   `gorm:"column:price" json:"price"`
	Point       int       `gorm:"column:point" json:"point"`
	Description string    `gorm:"column:description" json:"description"`
	Status      bool      `gorm:"column:status;default:true" json:"status"`
	PhoneNumber *string   `gorm:"column:phone_number" json:"phone_number"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
