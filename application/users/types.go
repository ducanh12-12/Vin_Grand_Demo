package users

import (
	"base-go/domain/model"
)

// command
type AddUserIpt struct {
	FullName    *string `json:"fullname"  validate:"required"`
	UserName    string  `json:"username"  validate:"required"`
	PhoneNumber *string `json:"phone_number"`
	Avatar      *string `json:"avatar" `
	Password    *string `json:"password" `
}
type UpdateUserIpt struct {
	FullName    *string `json:"fullname"`
	PhoneNumber *string `json:"phone_number"`
	Avatar      *string `json:"avatar" `
	Password    *string `json:"password" `
}
type AddPoint struct {
	InvoiceId int `json:"invoice_id" validate:"required"`
}

// query

// response
type GetUserResp = model.User
