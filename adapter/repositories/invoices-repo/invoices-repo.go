package invoices_repo

import (
	"base-go/domain/model"
	implFor "base-go/services/requires"
	"context"

	"gorm.io/gorm"
)

type invoicesRepo struct {
	db *gorm.DB
}

func NewInvoicesRepo(db *gorm.DB) implFor.InvoiceRepository {
	return &invoicesRepo{db: db}
}

func (r *invoicesRepo) Create(ctx context.Context, invoice model.Invoice) (*model.Invoice, error) {
	err := r.db.Create(&invoice).Error
	return &invoice, err
}
func (r *invoicesRepo) Update(ctx context.Context, id int, invoice model.Invoice) (*model.Invoice, error) {
	eventOld := model.Invoice{}
	if err := r.db.Where("id = ?", id).First(&eventOld).Error; err != nil {
		return &eventOld, err
	}
	eventCopy := invoice
	err := r.db.Model(&eventOld).Updates(&eventCopy).Error
	return &eventOld, err
}

func (r *invoicesRepo) Retrieve(ctx context.Context, id int) (*model.Invoice, error) {
	invoice := model.Invoice{}
	err := r.db.Where("id = ?", id).First(&invoice).Error
	return &invoice, err
}
func (r *invoicesRepo) List(ctx context.Context) (*[]model.Invoice, error) {
	var invoice []model.Invoice
	err := r.db.Find(&invoice).Error
	return &invoice, err
}
func (r *invoicesRepo) Delete(ctx context.Context, id int) (string, error) {
	err := r.db.Delete(model.Invoice{}, id).Error
	return "okla", err
}
