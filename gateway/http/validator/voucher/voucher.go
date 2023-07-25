package voucher

import (
	"base-go/application/vouchers"

	"github.com/go-playground/validator"
)

func ValidateVoucher(voucher vouchers.AddVoucherIpt) error {
	type Err struct {
		name string
	}
	validate := validator.New()

	err := validate.Struct(voucher)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// for _, fieldErr := range validationErrors {
		// 	fieldErr.Value() = fmt.Sprintf(fieldErr)
		// }
		return validationErrors
	}
	return nil
}
