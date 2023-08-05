package model

import "time"

type Voucher struct {
	Id        int        `gorm:"column:id;primaryKey;autoIncrement" json:"id" validate:"required"`
	Title     string     `gorm:"column:title" json:"title"`
	EventId   int        `gorm:"column:event_id" json:"event_id"`
	Point     int        `gorm:"column:point" json:"point"`
	Quantity  int        `gorm:"column:quantity" json:"quantity"`
	User      []User     `gorm:"many2many:user_vouchers;"`
	StartTime *time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime   *time.Time `gorm:"column:end_time" json:"end_time"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
}
