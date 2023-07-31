package invoices

import (
	"base-go/domain/model"
)

// command
type AddInvoiceIpt struct {
	Address     *string `json:"address" validate:"required"`
	Price       float64 `json:"price"  validate:"required"`
	Point       int     `json:"point"`
	Description string  `json:"description" `
	PhoneNumber *string `json:"phone_number" validate:"required,phone_number"`
	Status      bool    `json:"status"`
}
type UpdateInvoiceIpt struct {
	Address     *string `json:"address" validate:"required"`
	Price       float64 `json:"price"  validate:"required"`
	Point       int     `json:"point"`
	Description string  `json:"description" `
	PhoneNumber *string `json:"phone_number" validate:"required,phone_number"`
	Status      bool    `json:"status"`
}

// query

// response
type GetInvoiceResp = model.Invoice
