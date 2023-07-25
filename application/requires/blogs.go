package requires

import (
	"base-go/domain/model"
	"context"
)

type BlogsService interface {
	CreateBlog(ctx context.Context, blog model.Blog) (*model.Blog, error)
	Retrieve(ctx context.Context, id int) (*model.Blog, error)
	GetBlogs(ctx context.Context) (*[]model.Blog, error)
}
