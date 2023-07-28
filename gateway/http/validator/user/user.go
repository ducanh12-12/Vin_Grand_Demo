package user

import (
	"base-go/application/users"
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

func isValidUsername(username string) bool {
	regex := regexp.MustCompile(`^[\p{L}\p{N}._-]+$`)
	return regex.MatchString(username)
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
	if !isValidUsername(user.UserName) {
		return fmt.Errorf("Invalid username")
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
