package requires

import (
	"base-go/domain/model"
	"context"
)

type UsersService interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	AddPointUser(ctx context.Context, user model.User) (*model.User, error)
	Retrieve(ctx context.Context, id int) (*model.User, error)
	GetUsers(ctx context.Context) (*[]model.User, error)
	GetUserByUserName(ctx context.Context, username string) (*model.User, error)
}
