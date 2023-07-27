package user

import (
	"base-go/application/users"
	"strings"

	"github.com/go-playground/validator"
)

type MultiError []error

func (me MultiError) Error() string {
	var messages []string
	for _, err := range me {
		messages = append(messages, err.Error())
	}
	return strings.Join(messages, "; ")
}
func ValidateUser(user users.AddUserIpt) error {

	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// for _, fieldErr := range validationErrors {
		// 	fieldErr.Value() = fmt.Sprintf(fieldErr)
		// }
		// for _, err := range validationErrors {

		// 	error = fmt.Errorf("%w", fmt.Sprintf("Lỗi kiểm tra dữ liệu cho trường %s: %s", err.Field(), err.Tag()))

		// }
		return validationErrors
	}
	return nil
}
func ValidatePoint(user users.AddPoint) error {

	validate := validator.New()

	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// for _, fieldErr := range validationErrors {
		// 	fieldErr.Value() = fmt.Sprintf(fieldErr)
		// }
		// for _, err := range validationErrors {

		// 	error = fmt.Errorf("%w", fmt.Sprintf("Lỗi kiểm tra dữ liệu cho trường %s: %s", err.Field(), err.Tag()))

		// }
		return validationErrors
	}
	return nil
}
