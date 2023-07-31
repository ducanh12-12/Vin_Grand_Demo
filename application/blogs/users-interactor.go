package blogs

import (
	"base-go/application/requires"
	"base-go/common/logger"
	"base-go/domain/model"
	"context"
	"time"
)

type BlogsInteractor struct {
	blogService requires.BlogsService
}

func NewBlogsInteractor(blogService requires.BlogsService) BlogsInteractor {
	return BlogsInteractor{blogService}
}
func (interactor *BlogsInteractor) CreateBlog(ctx context.Context, blog AddBlogIpt) (*model.Blog, error) {
	now := time.Now()
	newBlog := model.Blog{
		Title:       blog.Title,
		Content:     blog.Content,
		Description: blog.Description,
		Avatar:      blog.Avatar,
		Status:      blog.Status,
		CreatedAt:   now,
	}
	blogNew, err := interactor.blogService.CreateBlog(ctx, newBlog)
	if err != nil {
		logger.Error("Unable to add blog, error: %s", err.Error())
		return nil, err
	}
	return blogNew, nil
}
func (interactor *BlogsInteractor) Update(ctx context.Context, blog UpdateBlogIpt, id int) (*model.Blog, error) {
	now := time.Now()
	newBlog := model.Blog{
		Title:       blog.Title,
		Content:     blog.Content,
		Description: blog.Description,
		Avatar:      blog.Avatar,
		Status:      blog.Status,
		UpdatedAt:   now,
	}
	blogNew, err := interactor.blogService.Update(ctx, newBlog, id)
	if err != nil {
		logger.Error("Unable to add blog, error: %s", err.Error())
		return nil, err
	}
	return blogNew, nil
}
func (interactor *BlogsInteractor) GetBlog(ctx context.Context, id int) (*GetBlogResp, error) {
	return interactor.blogService.Retrieve(ctx, id)
}
func (interactor *BlogsInteractor) Delete(ctx context.Context, id int) (string, error) {
	return interactor.blogService.Delete(ctx, id)
}
func (interactor *BlogsInteractor) GetBlogs(ctx context.Context) (*[]GetBlogResp, error) {
	return interactor.blogService.GetBlogs(ctx)
}
