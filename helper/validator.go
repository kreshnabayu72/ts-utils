package helper

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// ValidateStruct is a helper function to validate a struct.
func ValidateStruct(v any) error {
	if validate == nil {
		validate = validator.New(validator.WithRequiredStructEnabled())
	}
	return validate.Struct(v)
}
