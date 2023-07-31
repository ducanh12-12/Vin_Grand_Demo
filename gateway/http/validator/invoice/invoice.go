package invoice

import (
	"base-go/application/invoices"
	"regexp"

	"github.com/go-playground/validator"
)

func phoneValidator(fl validator.FieldLevel) bool {
	// Định nghĩa biểu thức chính quy để xác thực số điện thoại
	phoneRegex := `(0([35789]|2[1-9]))([0-9]{8})\b`

	// Lấy giá trị của trường số điện thoại
	phoneNumber := fl.Field().String()

	// Sử dụng regex để kiểm tra số điện thoại hợp lệ
	match, _ := regexp.MatchString(phoneRegex, phoneNumber)

	return match
}
func ValidateInvoice(invoice invoices.AddInvoiceIpt) error {
	type Err struct {
		name string
	}
	validate := validator.New()
	validate.RegisterValidation("phone_number", phoneValidator)

	err := validate.Struct(invoice)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// for _, fieldErr := range validationErrors {
		// 	fieldErr.Value() = fmt.Sprintf(fieldErr)
		// }
		return validationErrors
	}
	return nil
}
func ValidateUpdateInvoice(invoice invoices.UpdateInvoiceIpt) error {
	type Err struct {
		name string
	}
	validate := validator.New()
	validate.RegisterValidation("phone_number", phoneValidator)

	err := validate.Struct(invoice)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// for _, fieldErr := range validationErrors {
		// 	fieldErr.Value() = fmt.Sprintf(fieldErr)
		// }
		return validationErrors
	}
	return nil
}
