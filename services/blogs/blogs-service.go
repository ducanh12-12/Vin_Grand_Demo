package blog

import (
	implFor "base-go/application/requires"
	"base-go/domain/model"
	"base-go/services/requires"
	"context"
)

type blogsServiceImpl struct {
	blogRepo requires.BlogRepository
}

func NewBlogService(blogRepo requires.BlogRepository) implFor.BlogsService {
	return &blogsServiceImpl{blogRepo}
}
func (svc *blogsServiceImpl) CreateBlog(ctx context.Context, blog model.Blog) (*model.Blog, error) {
	return svc.blogRepo.CreateBlog(ctx, blog)
}
func (svc *blogsServiceImpl) Update(ctx context.Context, blog model.Blog, id int) (*model.Blog, error) {
	return svc.blogRepo.Update(ctx, blog, id)
}
func (svc *blogsServiceImpl) GetBlogs(ctx context.Context) (*[]model.Blog, error) {
	return svc.blogRepo.GetBlogs(ctx)
}
func (svc *blogsServiceImpl) Retrieve(ctx context.Context, id int) (*model.Blog, error) {
	return svc.blogRepo.Retrieve(ctx, id)
}
func (svc *blogsServiceImpl) Delete(ctx context.Context, id int) (string, error) {
	return svc.blogRepo.Delete(ctx, id)
}
