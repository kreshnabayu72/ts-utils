package validation

import (
	"gitea.qwertysystem.net/BETS/ts-utils/v2/entity"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/parser"
	"gitea.qwertysystem.net/BETS/ts-utils/v2/validation/v2/custom_validator"
	"github.com/gofiber/fiber/v2"
	"github.com/kittipat1413/go-common/framework/validator"
)

// EXAMPLE USAGE:
// body, err := validation.ParseAndValidateBody[dto.MyDTO](ctx)
// OR
// query, err := validation.ParseAndValidateQuery[dto.MyQueryDTO](ctx)

// if err != nil {
// 	return err.SendResponse(ctx)
// }

func RawValidate[T any](data T) *entity.HttpError {
	v, validatorErr := validator.NewValidator(
		validator.WithTagNameFunc(validator.JSONTagNameFunc),
		validator.WithCustomValidator(new(custom_validator.NotBlankValidator)),
	)

	// Return if there is an error when creating the validator
	if validatorErr != nil {
		return entity.InternalServerError(validatorErr.Error())
	}

	// Validate the data and return if there is an error
	if validateErr := v.ValidateStruct(data); validateErr != nil {
		return entity.BadRequest(validateErr.Error())
	}

	return nil
}

// Compacted common validation to reduce repeated code.
// Intended to be used in controller
func ParseAndValidateBody[T any](ctx *fiber.Ctx) (*T, *entity.HttpError) {
	// Parse the body
	var body T
	if err := ctx.BodyParser(&body); err != nil {
		return nil, entity.InternalServerError(err.Error())
	}

	// Validate the body and return error if there is any
	validationErr := RawValidate(body)
	if validationErr != nil {
		return nil, validationErr
	}

	// Return the parsed and validated body
	return &body, validationErr
}

func ParseAndValidateQuery[T any](ctx *fiber.Ctx) (*T, *entity.HttpError) {
	// Parse the query
	rawQuery := ctx.Queries()
	parsedQuery, err := parser.ParseQuery[T](rawQuery)
	// Return if there is an error when parsing the query
	if err != nil {
		return nil, entity.BadRequest(err.Error())
	}

	// Validate the query and return if there is an error
	validationErr := RawValidate(parsedQuery)
	if validationErr != nil {
		return nil, validationErr
	}

	// Return the parsed and validated query
	return parsedQuery, nil
}
