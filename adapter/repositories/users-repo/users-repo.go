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
func (r *usersRepo) GetUsers(ctx context.Context) (*[]model.User, error) {
	var user []model.User
	err := r.db.Find(&user).Error
	return &user, err
}
