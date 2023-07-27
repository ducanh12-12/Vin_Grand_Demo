package users

import (
	"base-go/domain/model"
)

// command
type AddUserIpt struct {
	FullName string `json:"fullname"  validate:"required"`
	UserName string `json:"username"  validate:"required"`
	Avatar   string `json:"avatar" `
	Password string `json:"password"  validate:"required"`
}
type AddPoint struct {
	Point int `json:"point"  validate:"required"`
}

// query

// response
type GetUserResp = model.User
