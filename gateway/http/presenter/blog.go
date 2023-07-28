package presenter

import (
	"base-go/domain/model"
)

func Blog(blog *model.Blog) model.Blog {
	return model.Blog{
		Id:          blog.Id,
		Title:       blog.Title,
		Content:     blog.Content,
		Description: blog.Description,
		Avatar:      blog.Avatar,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	}
}
