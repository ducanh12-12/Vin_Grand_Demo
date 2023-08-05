package auth

import "fmt"

func ValidateLogin(password string, username string, phone_number string) error {
	if username == "" && phone_number == "" {
		return fmt.Errorf("Bad Request")
	}
	if password == "" {
		return fmt.Errorf("Password is required")
	}
	return nil
}
