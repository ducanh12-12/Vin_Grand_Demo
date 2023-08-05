package voucher

import (
	"base-go/application/vouchers"
	"fmt"
	"time"

	"github.com/go-playground/validator"
)

func isEndTimeAfterStartTime(startTimeStr string, endTimeStr string) (bool, error) {
	layout := "2/1/2006"
	startTime, err := time.Parse(layout, startTimeStr)
	if err != nil {
		return false, err
	}

	endTime, err := time.Parse(layout, endTimeStr)
	if err != nil {
		return false, err
	}
	return startTime.After(endTime), nil
}
func ValidateVoucher(voucher vouchers.AddVoucherIpt) error {
	type Err struct {
		name string
	}
	validate := validator.New()

	err := validate.Struct(voucher)
	if voucher.StartTime != "" && voucher.EndTime != "" {
		isValid, _ := isEndTimeAfterStartTime(voucher.StartTime, voucher.EndTime)
		if isValid {
			return fmt.Errorf("End time must be greater than start time")
		}
	}
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// for _, fieldErr := range validationErrors {
		// 	fieldErr.Value() = fmt.Sprintf(fieldErr)
		// }
		return validationErrors
	}
	return nil
}
func ValidateUpdateVoucher(voucher vouchers.UpdateVoucherIpt) error {
	type Err struct {
		name string
	}
	validate := validator.New()

	err := validate.Struct(voucher)
	if voucher.StartTime != "" && voucher.EndTime != "" {
		isValid, _ := isEndTimeAfterStartTime(voucher.StartTime, voucher.EndTime)
		if isValid {
			return fmt.Errorf("End time must be greater than start time")
		}
	}
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// for _, fieldErr := range validationErrors {
		// 	fieldErr.Value() = fmt.Sprintf(fieldErr)
		// }
		return validationErrors
	}
	return nil
}
