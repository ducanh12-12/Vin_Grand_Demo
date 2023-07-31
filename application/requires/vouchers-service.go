package requires

import (
	"base-go/domain/model"
	"context"
)

type VouchersService interface {
	CreateVoucher(ctx context.Context, voucher model.Voucher) (*model.Voucher, error)
	Retrieve(ctx context.Context, id int) (*model.Voucher, error)
	GetVouchers(ctx context.Context) (*[]model.Voucher, error)
	Update(ctx context.Context, voucher model.Voucher, id int) (*model.Voucher, error)
	Delete(ctx context.Context, id int) (string, error)
}
