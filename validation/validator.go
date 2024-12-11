package validation

import (
	"fmt"
	"mime/multipart"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Validator interface {
	ValidateFile(file *multipart.FileHeader, allowedTypes []string, maxSize int64) error
	ValidateStruct(c *fiber.Ctx, payload interface{}) error
}

type DefaultValidator struct{}

var validate *validator.Validate

func init() {
	validate = validator.New()
	RegisterCustomValidators(validate)
}

// RegisterCustomValidators registers custom validators
func RegisterCustomValidators(validate *validator.Validate) {
	validate.RegisterValidation("mongodb", func(fl validator.FieldLevel) bool {
		if fl.Field().Kind() == reflect.String {
			// Coba parse menjadi ObjectID
			_, err := primitive.ObjectIDFromHex(fl.Field().String())
			return err == nil
		}
		return false
	})
}

// ValidateStruct parses and validates the payload
func (v *DefaultValidator) ValidateStruct(c *fiber.Ctx, payload interface{}) error {
	// Parse the request body
	if err := c.BodyParser(payload); err != nil {
		return fmt.Errorf("invalid request payload: %w", err)
	}

	// Validate the struct
	if err := validate.Struct(payload); err != nil {
		return fmt.Errorf("validation errors: %s", FormatValidationError(err))
	}

	return nil
}

// FormatValidationError formats validation errors
func FormatValidationError(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		var formattedErrors []string
		for _, fieldError := range validationErrors {
			formattedErrors = append(formattedErrors, fieldError.Error())
		}
		return fmt.Sprintf("%v", formattedErrors)
	}
	return err.Error()
}
