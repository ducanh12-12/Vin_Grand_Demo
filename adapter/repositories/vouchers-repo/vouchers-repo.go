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
func (r *vouchersRepo) GetVoucherByTitle(ctx context.Context, title string) (*model.Voucher, error) {
	voucher := model.Voucher{}
	err := r.db.Where("title = ?", title).First(&voucher).Error
	return &voucher, err
}
func (r *vouchersRepo) GetVouchers(ctx context.Context) (*[]model.Voucher, error) {
	var voucher []model.Voucher
	err := r.db.Order("id asc").Find(&voucher).Error
	return &voucher, err
}
func (r *vouchersRepo) Update(ctx context.Context, voucher model.Voucher, id int) (*model.Voucher, error) {
	eventOld := model.Voucher{}
	if err := r.db.Where("id = ?", id).First(&eventOld).Error; err != nil {
		return &eventOld, err
	}

	err := r.db.Model(&eventOld).Updates(&voucher).Error
	return &eventOld, err
}
func (r *vouchersRepo) Delete(ctx context.Context, id int) (string, error) {
	err := r.db.Delete(model.Voucher{}, id).Error
	return "okla", err
}
