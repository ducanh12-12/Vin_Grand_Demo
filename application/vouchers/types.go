package vouchers

import (
	"base-go/domain/model"
	"time"
)

// command
type AddVoucherIpt struct {
	Title    string    `json:"title"  validate:"required"`
	EventId  int       `json:"eventId"`
	Point    int       `json:"point"  validate:"required"`
	Quantity int       `json:"quantity"  validate:"required"`
	OutDate  time.Time `json:"outdate"`
}
type UpdateVoucherIpt struct {
	Title    string    `json:"title"  validate:"required"`
	EventId  int       `json:"eventId"`
	Point    int       `json:"point"  validate:"required"`
	Quantity int       `json:"quantity"  validate:"required"`
	OutDate  time.Time `json:"outdate"`
}

// query

// response
type GetVoucherResp = model.Voucher
