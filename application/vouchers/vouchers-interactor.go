package vouchers

import (
	"base-go/application/requires"
	"base-go/common/logger"
	"base-go/domain/model"
	"context"
	"time"
)

type VouchersInteractor struct {
	voucherService requires.VouchersService
}

func NewVouchersInteractor(voucherService requires.VouchersService) VouchersInteractor {
	return VouchersInteractor{voucherService}
}
func (interactor *VouchersInteractor) CreateVoucher(ctx context.Context, voucher AddVoucherIpt) (*model.Voucher, error) {
	now := time.Now()
	newVoucher := model.Voucher{
		Title:     voucher.Title,
		EventId:   voucher.EventId,
		Point:     voucher.Point,
		Quantity:  voucher.Quantity,
		OutDate:   voucher.OutDate,
		CreatedAt: now,
	}
	voucherNew, err := interactor.voucherService.CreateVoucher(ctx, newVoucher)
	if err != nil {
		logger.Error("Unable to add voucher, error: %s", err.Error())
		return nil, err
	}
	return voucherNew, nil
}
func (interactor *VouchersInteractor) GetVoucher(ctx context.Context, id int) (*GetVoucherResp, error) {
	return interactor.voucherService.Retrieve(ctx, id)
}
func (interactor *VouchersInteractor) GetVouchers(ctx context.Context) (*[]GetVoucherResp, error) {
	return interactor.voucherService.GetVouchers(ctx)
}
