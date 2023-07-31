package model

import "time"

type Voucher struct {
	Id        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id" validate:"required"`
	Title     string    `gorm:"column:title" json:"title"`
	EventId   int       `gorm:"column:event_id" json:"event_id"`
	Point     int       `gorm:"column:point" json:"point"`
	Quantity  int       `gorm:"column:quantity" json:"quantity"`
	OutDate   time.Time `gorm:"column:out_date" json:"out_date"`
	User      []User    `gorm:"many2many:user_vouchers;"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
