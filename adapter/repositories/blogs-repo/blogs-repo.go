package blogs_repo

import (
	"base-go/domain/model"
	implFor "base-go/services/requires"
	"context"

	"gorm.io/gorm"
)

type blogsRepo struct {
	db *gorm.DB
}

func NewBlogsRepo(db *gorm.DB) implFor.BlogRepository {
	return &blogsRepo{db: db}
}

func (r *blogsRepo) CreateBlog(ctx context.Context, blog model.Blog) (*model.Blog, error) {
	err := r.db.Create(&blog).Error
	return &blog, err
}

func (r *blogsRepo) Retrieve(ctx context.Context, id int) (*model.Blog, error) {
	blog := model.Blog{}
	err := r.db.Where("id = ?", id).First(&blog).Error
	return &blog, err
}
func (r *blogsRepo) GetBlogs(ctx context.Context) (*[]model.Blog, error) {
	var blog []model.Blog
	err := r.db.Find(&blog).Error
	return &blog, err
}
