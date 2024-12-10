package validation

import (
	"github.com/betsv2/tsv2-go-utils/entity"
	"github.com/betsv2/tsv2-go-utils/validation/v2/custom_validator"
	"github.com/gofiber/fiber/v2"
	"github.com/kittipat1413/go-common/framework/validator"
)

// EXAMPLE USAGE:
// body, validationErr := validation.MakeAndValidate[dto.MyDTO](ctx)

// if validationErr != nil {
// 	return response.SendResponse(ctx, validationErr.Code, nil, validationErr.Message)
// }

// Compacted common validation to reduce repeated code.
// Intended to be used in controller
func MakeAndValidate[T any](ctx *fiber.Ctx) (T, *entity.HttpError) {
	var body T
	if err := ctx.BodyParser(&body); err != nil {
		return body, entity.InternalServerError(err.Error())
	}

	// Register the validator
	v, validatorErr := validator.NewValidator(
		validator.WithTagNameFunc(validator.JSONTagNameFunc),
		validator.WithCustomValidator(new(custom_validator.NotBlankValidator)),
	)

	if validatorErr != nil {
		return body, entity.InternalServerError(validatorErr.Error())
	}

	// Validate the body
	if validateErr := v.ValidateStruct(body); validateErr != nil {
		return body, entity.BadRequest(validateErr.Error())
	}

	return body, nil
}
