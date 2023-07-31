package invoices

import (
	"base-go/application/requires"
	"base-go/common/logger"
	"base-go/domain/model"
	"context"
	"time"
)

type InvoicesInteractor struct {
	invoiceService requires.InvoicesService
}

func NewInvoicesInteractor(invoiceService requires.InvoicesService) InvoicesInteractor {
	return InvoicesInteractor{invoiceService}
}
func (interactor *InvoicesInteractor) Create(ctx context.Context, invoice AddInvoiceIpt) (*model.Invoice, error) {
	now := time.Now()
	newInvoice := model.Invoice{
		Address:     invoice.Address,
		Price:       invoice.Price,
		Point:       invoice.Point,
		Description: invoice.Description,
		Status:      invoice.Status,
		PhoneNumber: invoice.PhoneNumber,
		CreatedAt:   now,
	}
	invoiceNew, err := interactor.invoiceService.Create(ctx, newInvoice)
	if err != nil {
		logger.Error("Unable to add invoice, error: %s", err.Error())
		return nil, err
	}
	return invoiceNew, nil
}
func (interactor *InvoicesInteractor) Update(ctx context.Context, id int, invoice UpdateInvoiceIpt) (*model.Invoice, error) {
	now := time.Now()
	updateInvoice := model.Invoice{
		Address:     invoice.Address,
		Price:       invoice.Price,
		Point:       invoice.Point,
		Description: invoice.Description,
		Status:      invoice.Status,
		PhoneNumber: invoice.PhoneNumber,
		UpdatedAt:   now,
	}
	invoiceNew, err := interactor.invoiceService.Update(ctx, id, updateInvoice)
	if err != nil {
		logger.Error("Unable to add invoice, error: %s", err.Error())
		return nil, err
	}
	return invoiceNew, nil
}
func (interactor *InvoicesInteractor) Retrieve(ctx context.Context, id int) (*GetInvoiceResp, error) {
	return interactor.invoiceService.Retrieve(ctx, id)
}
func (interactor *InvoicesInteractor) Delete(ctx context.Context, id int) (string, error) {
	return interactor.invoiceService.Delete(ctx, id)
}
func (interactor *InvoicesInteractor) List(ctx context.Context) (*[]GetInvoiceResp, error) {
	return interactor.invoiceService.List(ctx)
}
