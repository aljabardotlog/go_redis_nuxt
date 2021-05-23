package handlers

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func Nama(field validator.FieldLevel) bool {
	if field.Field().String() != "" {
		valid := regexp.MustCompile(`^([^0\+]{1})[\d]+$`)
		return valid.MatchString(field.Field().String())
	}

	return true
}

func NomorTelepon(field validator.FieldLevel) bool {
	if field.Field().String() != "" {
		valid := regexp.MustCompile(`^([^0\+]{1})[\d]+$`)
		return valid.MatchString(field.Field().String())
	}

	return true
}

func DateTime(field validator.FieldLevel) bool {
	if field.Field().String() != "" {
		valid := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2}) (\d{2}):(\d{2}):(\d{2})`)
		return valid.MatchString(field.Field().String())
	}

	return true
}

func ValidateUsername(field validator.FieldLevel) bool {
	usernameRegexp := regexp.MustCompile(`^[[:alnum:]_]+$`)
	return usernameRegexp.MatchString(field.Field().String())
}
