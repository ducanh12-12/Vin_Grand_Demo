package requires

import (
	"base-go/domain/model"
	"context"
)

type BlogRepository interface {
	CreateBlog(ctx context.Context, blog model.Blog) (*model.Blog, error)
	Update(ctx context.Context, blog model.Blog, id int) (*model.Blog, error)
	Retrieve(ctx context.Context, id int) (*model.Blog, error)
	GetBlogs(ctx context.Context) (*[]model.Blog, error)
	Delete(ctx context.Context, id int) (string, error)
}
