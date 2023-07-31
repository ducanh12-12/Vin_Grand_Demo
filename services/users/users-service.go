package users

import (
	implFor "base-go/application/requires"
	"base-go/domain/model"
	"base-go/services/requires"
	"context"
)

type usersServiceImpl struct {
	userRepo requires.UserRepository
}

func NewUserService(userRepo requires.UserRepository) implFor.UsersService {
	return &usersServiceImpl{userRepo}
}
func (svc *usersServiceImpl) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	return svc.userRepo.CreateUser(ctx, user)
}
func (svc *usersServiceImpl) Update(ctx context.Context, user model.User, id int) (*model.User, error) {
	return svc.userRepo.Update(ctx, user, id)
}
func (svc *usersServiceImpl) AddPointUser(ctx context.Context, user model.User) (*model.User, error) {
	return svc.userRepo.AddPointUser(ctx, user)
}

func (svc *usersServiceImpl) GetUsers(ctx context.Context) (*[]model.User, error) {
	return svc.userRepo.GetUsers(ctx)
}
func (svc *usersServiceImpl) GetUserByUserName(ctx context.Context, username string) (*model.User, error) {
	return svc.userRepo.GetUserByUserName(ctx, username)
}
func (svc *usersServiceImpl) GetUserByPhoneNumber(ctx context.Context, phone_number string) (*model.User, error) {
	return svc.userRepo.GetUserByPhoneNumber(ctx, phone_number)
}

func (svc *usersServiceImpl) Retrieve(ctx context.Context, id int) (*model.User, error) {
	return svc.userRepo.Retrieve(ctx, id)
}
func (svc *usersServiceImpl) Delete(ctx context.Context, id int) (string, error) {
	return svc.userRepo.Delete(ctx, id)
}
