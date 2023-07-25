package vouchers

import (
	implFor "base-go/application/requires"
	"base-go/domain/model"
	"base-go/services/requires"
	"context"
)

type vouchersServiceImpl struct {
	voucherRepo requires.VoucherRepository
}

func NewVoucherService(voucherRepo requires.VoucherRepository) implFor.VouchersService {
	return &vouchersServiceImpl{voucherRepo}
}
func (svc *vouchersServiceImpl) CreateVoucher(ctx context.Context, voucher model.Voucher) (*model.Voucher, error) {
	return svc.voucherRepo.CreateVoucher(ctx, voucher)
}

func (svc *vouchersServiceImpl) GetVouchers(ctx context.Context) (*[]model.Voucher, error) {
	return svc.voucherRepo.GetVouchers(ctx)
}
func (svc *vouchersServiceImpl) Retrieve(ctx context.Context, id int) (*model.Voucher, error) {
	return svc.voucherRepo.Retrieve(ctx, id)
}
