package requires

import (
	"base-go/domain/model"
	"context"
)

type InvoicesService interface {
	Create(ctx context.Context, invoice model.Invoice) (*model.Invoice, error)
	Update(ctx context.Context, id int, invoice model.Invoice) (*model.Invoice, error)
	Retrieve(ctx context.Context, id int) (*model.Invoice, error)
	List(ctx context.Context) (*[]model.Invoice, error)
	Delete(ctx context.Context, id int) (string, error)
}
