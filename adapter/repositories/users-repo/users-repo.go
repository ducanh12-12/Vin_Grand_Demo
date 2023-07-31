package cats_repo

import (
	"base-go/domain/model"
	implFor "base-go/services/requires"
	"context"

	"gorm.io/gorm"
)

type usersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) implFor.UserRepository {
	return &usersRepo{db: db}
}

func (r *usersRepo) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	err := r.db.Create(&user).Error
	return &user, err
}
func (r *usersRepo) AddPointUser(ctx context.Context, user model.User) (*model.User, error) {
	err := r.db.Where("id = ?", user.Id).Updates(user).First(&user).Error
	return &user, err
}

func (r *usersRepo) Retrieve(ctx context.Context, id int) (*model.User, error) {
	user := model.User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}
func (r *usersRepo) GetUserByUserName(ctx context.Context, username string) (*model.User, error) {
	user := model.User{}
	err := r.db.Where("user_name = ?", username).First(&user).Error
	return &user, err
}
func (r *usersRepo) GetUserByPhoneNumber(ctx context.Context, phone_number string) (*model.User, error) {
	user := model.User{}
	err := r.db.Where("phone_number = ?", phone_number).First(&user).Error
	return &user, err
}
func (r *usersRepo) GetUsers(ctx context.Context) (*[]model.User, error) {
	var user []model.User
	err := r.db.Order("id asc").Find(&user).Error
	return &user, err
}
func (r *usersRepo) Update(ctx context.Context, user model.User, id int) (*model.User, error) {
	eventOld := model.User{}
	if err := r.db.Where("id = ?", id).First(&eventOld).Error; err != nil {
		return &eventOld, err
	}
	userCopy := user
	err := r.db.Model(&eventOld).Updates(&userCopy).Error
	return &eventOld, err
}
func (r *usersRepo) Delete(ctx context.Context, id int) (string, error) {
	err := r.db.Delete(model.User{}, id).Error
	return "okla", err
}
