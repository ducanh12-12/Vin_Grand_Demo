package event

import (
	"base-go/application/events"
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
func ValidateEvent(event events.AddEventIpt) error {
	type Err struct {
		name string
	}
	validate := validator.New()

	err := validate.Struct(event)
	if event.StartTime != "" && event.EndTime != "" {
		isValid, _ := isEndTimeAfterStartTime(event.StartTime, event.EndTime)
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
func ValidateUpdateEvent(event events.UpdateEventIpt) error {
	type Err struct {
		name string
	}
	validate := validator.New()
	if event.StartTime != "" && event.EndTime != "" {
		isValid, _ := isEndTimeAfterStartTime(event.StartTime, event.EndTime)
		if isValid {
			return fmt.Errorf("End time must be greater than start time")
		}
	}
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
