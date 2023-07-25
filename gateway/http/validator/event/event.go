package event

import (
	"base-go/application/events"

	"github.com/go-playground/validator"
)

func ValidateEvent(event events.AddEventIpt) error {
	type Err struct {
		name string
	}
	validate := validator.New()

	err := validate.Struct(event)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		// for _, fieldErr := range validationErrors {
		// 	fieldErr.Value() = fmt.Sprintf(fieldErr)
		// }
		return validationErrors
	}
	return nil
}
