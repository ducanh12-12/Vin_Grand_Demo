package model

import (
	"time"
)

type User struct {
	Id        int       `gorm:"column:id;primaryKey;autoIncrement" json:"id" validate:"required"`
	FullName  string    `gorm:"column:full_name" json:"fullname"`
	UserName  string    `gorm:"column:user_name" json:"username"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Password  string    `gorm:"column:password" json:"-"`
	LastLogin string    `gorm:"column:last_login" json:"last_login"`
	Point     int       `gorm:"column:point" json:"point"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
