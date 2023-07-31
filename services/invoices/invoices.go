package invoices

import (
	implFor "base-go/application/requires"
	"base-go/domain/model"
	"base-go/services/requires"
	"context"
)

type invoicesServiceImpl struct {
	invoiceRepo requires.InvoiceRepository
}

func NewInvoiceService(invoiceRepo requires.InvoiceRepository) implFor.InvoicesService {
	return &invoicesServiceImpl{invoiceRepo}
}
func (svc *invoicesServiceImpl) Create(ctx context.Context, invoice model.Invoice) (*model.Invoice, error) {
	return svc.invoiceRepo.Create(ctx, invoice)
}
func (svc *invoicesServiceImpl) Update(ctx context.Context, id int, invoice model.Invoice) (*model.Invoice, error) {
	return svc.invoiceRepo.Update(ctx, id, invoice)
}
func (svc *invoicesServiceImpl) List(ctx context.Context) (*[]model.Invoice, error) {
	return svc.invoiceRepo.List(ctx)
}
func (svc *invoicesServiceImpl) Retrieve(ctx context.Context, id int) (*model.Invoice, error) {
	return svc.invoiceRepo.Retrieve(ctx, id)
}
func (svc *invoicesServiceImpl) Delete(ctx context.Context, id int) (string, error) {
	return svc.invoiceRepo.Delete(ctx, id)
}
