package users

import (
	"base-go/application/requires"
	"base-go/common/logger"
	"base-go/domain/model"
	"context"
	"time"
)

type UsersInteractor struct {
	userService requires.UsersService
}

func NewUsersInteractor(userService requires.UsersService) UsersInteractor {
	return UsersInteractor{userService}
}
func (interactor *UsersInteractor) AddUser(ctx context.Context, user AddUserIpt) (*model.User, error) {
	now := time.Now()
	newUser := model.User{
		FullName:  user.FullName,
		UserName:  user.UserName,
		Avatar:    user.Avatar,
		Password:  user.Password,
		CreatedAt: now,
	}
	userNew, err := interactor.userService.CreateUser(ctx, newUser)
	if err != nil {
		logger.Error("Unable to add user, error: %s", err.Error())
		return nil, err
	}
	return userNew, nil
}
func (interactor *UsersInteractor) GetUser(ctx context.Context, id int) (*GetUserResp, error) {
	return interactor.userService.Retrieve(ctx, id)
}
func (interactor *UsersInteractor) GetUserByUserName(ctx context.Context, username string) (*GetUserResp, error) {
	return interactor.userService.GetUserByUserName(ctx, username)
}
func (interactor *UsersInteractor) GetUsers(ctx context.Context) (*[]GetUserResp, error) {
	return interactor.userService.GetUsers(ctx)
}
