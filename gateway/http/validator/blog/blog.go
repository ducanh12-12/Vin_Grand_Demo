package blog

import (
	"base-go/application/blogs"

	"github.com/go-playground/validator"
)

func ValidateBlog(blog blogs.AddBlogIpt) error {
	type Err struct {
		name string
	}
	validate := validator.New()

	err := validate.Struct(blog)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// for _, fieldErr := range validationErrors {
		// 	fieldErr.Value() = fmt.Sprintf(fieldErr)
		// }
		return validationErrors
	}
	return nil
}
