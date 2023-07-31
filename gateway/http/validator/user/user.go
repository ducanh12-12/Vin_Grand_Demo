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
func phoneValidator(phone_number string) bool {
	phoneRegex := `(0([35789]|2[1-9]))([0-9]{8})\b`

	// Sử dụng regex để kiểm tra số điện thoại hợp lệ
	match, _ := regexp.MatchString(phoneRegex, phone_number)

	return match
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
	if user.PhoneNumber != nil {
		if !phoneValidator(*user.PhoneNumber) {
			return fmt.Errorf("Invalid phone number")
		}
	}
	if !isValidUsername(user.UserName) {
		return fmt.Errorf("Invalid username")
	}
	return nil
}
func ValidateUpdateUser(user users.UpdateUserIpt) error {

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
	if user.PhoneNumber != nil {
		if !phoneValidator(*user.PhoneNumber) {
			return fmt.Errorf("Invalid phone number")
		}
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
