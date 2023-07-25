package presenter

import (
	"base-go/domain/model"
)

func Voucher(voucher *model.Voucher) model.Voucher {
	return model.Voucher{
		Title:     voucher.Title,
		EventId:   voucher.EventId,
		Point:     voucher.Point,
		Quantity:  voucher.Quantity,
		OutDate:   voucher.OutDate,
		UpdatedAt: voucher.UpdatedAt,
		CreatedAt: voucher.CreatedAt,
	}
}
