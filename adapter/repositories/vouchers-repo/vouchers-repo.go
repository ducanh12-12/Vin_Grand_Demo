package vouchers_repo

import (
	"base-go/domain/model"
	implFor "base-go/services/requires"
	"context"

	"gorm.io/gorm"
)

type vouchersRepo struct {
	db *gorm.DB
}

func NewVouchersRepo(db *gorm.DB) implFor.VoucherRepository {
	return &vouchersRepo{db: db}
}

func (r *vouchersRepo) CreateVoucher(ctx context.Context, voucher model.Voucher) (*model.Voucher, error) {
	err := r.db.Create(&voucher).Error
	return &voucher, err
}

func (r *vouchersRepo) Retrieve(ctx context.Context, id int) (*model.Voucher, error) {
	voucher := model.Voucher{}
	err := r.db.Where("id = ?", id).First(&voucher).Error
	return &voucher, err
}
func (r *vouchersRepo) GetVouchers(ctx context.Context) (*[]model.Voucher, error) {
	var voucher []model.Voucher
	err := r.db.Find(&voucher).Error
	return &voucher, err
}
