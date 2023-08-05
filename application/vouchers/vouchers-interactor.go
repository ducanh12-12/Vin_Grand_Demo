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
	layout := "2/1/2006"
	var startTime *time.Time
	var endTime *time.Time
	var err error
	parsedStartTime, err := time.Parse(layout, voucher.StartTime)
	if err == nil {
		startTime = &parsedStartTime
	}
	parsedEndTime, err := time.Parse(layout, voucher.EndTime)
	if err == nil {
		endTime = &parsedEndTime
	}
	newVoucher := model.Voucher{
		Title:     voucher.Title,
		EventId:   voucher.EventId,
		Point:     voucher.Point,
		Quantity:  voucher.Quantity,
		StartTime: startTime,
		EndTime:   endTime,
		CreatedAt: now,
	}
	voucherNew, err := interactor.voucherService.CreateVoucher(ctx, newVoucher)
	if err != nil {
		logger.Error("Unable to add voucher, error: %s", err.Error())
		return nil, err
	}
	return voucherNew, nil
}
func (interactor *VouchersInteractor) Update(ctx context.Context, voucher UpdateVoucherIpt, id int) (*model.Voucher, error) {
	now := time.Now()
	layout := "2/1/2006"
	var startTime *time.Time
	var endTime *time.Time
	var err error
	parsedStartTime, err := time.Parse(layout, voucher.StartTime)
	if err == nil {
		startTime = &parsedStartTime
	}
	parsedEndTime, err := time.Parse(layout, voucher.EndTime)
	if err == nil {
		endTime = &parsedEndTime
	}
	newVoucher := model.Voucher{
		Title:     voucher.Title,
		EventId:   voucher.EventId,
		Point:     voucher.Point,
		Quantity:  voucher.Quantity,
		StartTime: startTime,
		EndTime:   endTime,
		UpdatedAt: now,
	}
	voucherNew, err := interactor.voucherService.Update(ctx, newVoucher, id)
	if err != nil {
		logger.Error("Unable to add voucher, error: %s", err.Error())
		return nil, err
	}
	return voucherNew, nil
}
func (interactor *VouchersInteractor) GetVoucher(ctx context.Context, id int) (*GetVoucherResp, error) {
	return interactor.voucherService.Retrieve(ctx, id)
}
func (interactor *VouchersInteractor) GetVoucherByTitle(ctx context.Context, title string) (*GetVoucherResp, error) {
	return interactor.voucherService.GetVoucherByTitle(ctx, title)
}
func (interactor *VouchersInteractor) Delete(ctx context.Context, id int) (string, error) {
	return interactor.voucherService.Delete(ctx, id)
}
func (interactor *VouchersInteractor) GetVouchers(ctx context.Context) (*[]GetVoucherResp, error) {
	return interactor.voucherService.GetVouchers(ctx)
}
