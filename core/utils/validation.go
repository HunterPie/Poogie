package utils

import (
	"net/mail"
)

func ValidateEmail(value string) bool {
	_, err := mail.ParseAddress(value)

	return err == nil
}
