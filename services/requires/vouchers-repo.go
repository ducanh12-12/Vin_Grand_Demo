package requires

import (
	"base-go/domain/model"
	"context"
)

type VoucherRepository interface {
	CreateVoucher(ctx context.Context, voucher model.Voucher) (*model.Voucher, error)
	Update(ctx context.Context, voucher model.Voucher, id int) (*model.Voucher, error)
	Retrieve(ctx context.Context, id int) (*model.Voucher, error)
	GetVoucherByTitle(ctx context.Context, title string) (*model.Voucher, error)
	Delete(ctx context.Context, id int) (string, error)
	GetVouchers(ctx context.Context) (*[]model.Voucher, error)
}
