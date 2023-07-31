package presenter

import (
	"base-go/domain/model"
)

func User(user *model.User) model.User {
	return model.User{
		Id:          user.Id,
		FullName:    user.FullName,
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Point:       user.Point,
	}
}
