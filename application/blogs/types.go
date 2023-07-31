package blogs

import (
	"base-go/domain/model"
)

// command
type AddBlogIpt struct {
	Title       string `json:"title"  validate:"required"`
	Content     string `json:"content"  validate:"required"`
	Description string `json:"description"`
	Avatar      string `json:"avatar" `
	Status      bool   `json:"status"`
}
type UpdateBlogIpt struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
	Avatar      string `json:"avatar" `
	Status      bool   `json:"status"`
}

// query

// response
type GetBlogResp = model.Blog
