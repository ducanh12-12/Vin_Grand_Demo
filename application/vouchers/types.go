package vouchers

import (
	"base-go/domain/model"
)

// command
type AddVoucherIpt struct {
	Title     string `json:"title"  validate:"required"`
	EventId   int    `json:"event_id"`
	Point     int    `json:"point"  validate:"required"`
	Quantity  int    `json:"quantity"  validate:"required"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
type UpdateVoucherIpt struct {
	Title     string `json:"title"  validate:"required"`
	EventId   int    `json:"event_id"`
	Point     int    `json:"point"  validate:"required"`
	Quantity  int    `json:"quantity"  validate:"required"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

// query

// response
type GetVoucherResp = model.Voucher
