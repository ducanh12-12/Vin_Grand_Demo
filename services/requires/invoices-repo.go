package requires

import (
	"base-go/domain/model"
	"context"
)

type InvoiceRepository interface {
	Create(ctx context.Context, event model.Invoice) (*model.Invoice, error)
	Update(ctx context.Context, id int, event model.Invoice) (*model.Invoice, error)
	Retrieve(ctx context.Context, id int) (*model.Invoice, error)
	Delete(ctx context.Context, id int) (string, error)
	List(ctx context.Context) (*[]model.Invoice, error)
}
