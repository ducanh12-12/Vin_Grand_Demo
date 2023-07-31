package presenter

import (
	"base-go/domain/model"
)

func Invoice(invoice *model.Invoice) model.Invoice {
	return model.Invoice{
		Id:          invoice.Id,
		Address:     invoice.Address,
		Price:       invoice.Price,
		Description: invoice.Description,
		Point:       invoice.Point,
		Status:      invoice.Status,
		PhoneNumber: invoice.PhoneNumber,
		CreatedAt:   invoice.CreatedAt,
		UpdatedAt:   invoice.UpdatedAt,
	}
}
