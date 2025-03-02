package entity

import (
	"gitea.qwertysystem.net/BETS/ts-utils/response"
	"github.com/gofiber/fiber/v2"
)

type HttpError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Err     interface{} `json:"error,omitempty"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func (e *HttpError) SendResponse(ctx *fiber.Ctx) error {
	return response.SendResponse(ctx, e.Code, nil, e.Message)
}

func (e *HttpError) SendResponseWithError(ctx *fiber.Ctx) error {
	return response.SendResponseWithError(ctx, e.Code, nil, e.Message, e.Err)
}

func InternalServerError(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusInternalServerError,
		Message: message,
	}
}

func BadRequest(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusBadRequest,
		Message: message,
	}
}

func Unauthorized(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusUnauthorized,
		Message: message,
	}
}

func Forbidden(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusForbidden,
		Message: message,
	}
}

func NotFound(message string) *HttpError {
	return &HttpError{
		Code:    fiber.StatusNotFound,
		Message: message,
	}
}

func ToHttpError(err error) *HttpError {
	if httpErr, ok := err.(*HttpError); ok {
		return httpErr
	}

	return InternalServerError(err.Error())
}
