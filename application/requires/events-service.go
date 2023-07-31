package requires

import (
	"base-go/domain/model"
	"context"
)

type EventsService interface {
	CreateEvent(ctx context.Context, event model.Event) (*model.Event, error)
	Retrieve(ctx context.Context, id int) (*model.Event, error)
	GetEvents(ctx context.Context) (*[]model.Event, error)
	Update(ctx context.Context, event model.Event, id int) (*model.Event, error)
	Delete(ctx context.Context, id int) (string, error)
}
