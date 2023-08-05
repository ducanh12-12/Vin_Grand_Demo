package presenter

import (
	"base-go/domain/model"
)

func Voucher(voucher *model.Voucher) model.Voucher {
	return model.Voucher{
		Id:        voucher.Id,
		Title:     voucher.Title,
		EventId:   voucher.EventId,
		Point:     voucher.Point,
		Quantity:  voucher.Quantity,
		StartTime: voucher.StartTime,
		EndTime:   voucher.EndTime,
		UpdatedAt: voucher.UpdatedAt,
		CreatedAt: voucher.CreatedAt,
	}
}
