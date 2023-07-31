package users

import (
	"base-go/application/requires"
	"base-go/common/logger"
	"base-go/domain/model"
	"context"
	"fmt"
	"time"
)

type UsersInteractor struct {
	userService    requires.UsersService
	invoiceService requires.InvoicesService
}

func NewUsersInteractor(userService requires.UsersService, invoiceService requires.InvoicesService) UsersInteractor {
	return UsersInteractor{userService, invoiceService}
}
func (interactor *UsersInteractor) AddUser(ctx context.Context, user AddUserIpt) (*model.User, error) {
	now := time.Now()
	newUser := model.User{
		FullName:    user.FullName,
		UserName:    user.UserName,
		Avatar:      user.Avatar,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   now,
	}
	userNew, err := interactor.userService.CreateUser(ctx, newUser)
	if err != nil {
		logger.Error("Unable to add user, error: %s", err.Error())
		return nil, err
	}
	return userNew, nil
}
func (interactor *UsersInteractor) Update(ctx context.Context, user UpdateUserIpt, id int) (*model.User, error) {
	now := time.Now()
	newUser := model.User{
		FullName:    user.FullName,
		Avatar:      user.Avatar,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		UpdatedAt:   now,
	}

	userNew, err := interactor.userService.Update(ctx, newUser, id)
	if err != nil {
		logger.Error("Unable to add user, error: %s", err.Error())
		return nil, err
	}
	return userNew, nil
}
func (interactor *UsersInteractor) GetUser(ctx context.Context, id int) (*GetUserResp, error) {
	return interactor.userService.Retrieve(ctx, id)
}
func (interactor *UsersInteractor) Delete(ctx context.Context, id int) (string, error) {
	return interactor.userService.Delete(ctx, id)
}
func (interactor *UsersInteractor) GetUserByUserName(ctx context.Context, username string) (*GetUserResp, error) {
	return interactor.userService.GetUserByUserName(ctx, username)
}
func (interactor *UsersInteractor) GetUserByPhoneNumber(ctx context.Context, phone_number string) (*GetUserResp, error) {
	return interactor.userService.GetUserByPhoneNumber(ctx, phone_number)
}
func (interactor *UsersInteractor) GetUsers(ctx context.Context) (*[]GetUserResp, error) {
	return interactor.userService.GetUsers(ctx)
}
func (interactor *UsersInteractor) AddPointUser(ctx context.Context, invoice_id AddPoint, id int) (*model.User, error) {
	invoice, err := interactor.invoiceService.Retrieve(ctx, invoice_id.InvoiceId)
	if err != nil || invoice.Status == false {
		return nil, fmt.Errorf("Invalid invoice")
	}

	olduser, err := interactor.userService.Retrieve(ctx, id)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	point := model.User{
		Id:        olduser.Id,
		Point:     olduser.Point + invoice.Point,
		UpdatedAt: now,
	}
	pointNew, err := interactor.userService.AddPointUser(ctx, point)
	if err != nil {
		logger.Error("Unable to add user, error: %s", err.Error())
		return nil, err
	}
	invoice.Status = false
	interactor.invoiceService.Update(ctx, invoice.Id, *invoice)
	return pointNew, nil
}
