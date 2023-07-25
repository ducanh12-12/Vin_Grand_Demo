package vouchers

import (
	"base-go/domain/model"
	"time"
)

// command
type AddVoucherIpt struct {
	Title    string    `json:"title"  validate:"required"`
	EventId  int       `json:"eventId"  validate:"required"`
	Point    int       `json:"point"  validate:"required"`
	Quantity int       `json:"quantity"  validate:"required"`
	OutDate  time.Time `json:"outdate"  validate:"required"`
}

// query

// response
type GetVoucherResp = model.Voucher
