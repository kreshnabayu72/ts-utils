package functions

import "github.com/gofiber/fiber/v2"

type HttpResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Err     any    `json:"error,omitempty"`
}

// SendResponse is a helper to send JSON responses in Fiber
func SendResponse(c *fiber.Ctx, statusCode int, data any, message string, err any) error {
	response := HttpResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
		Err:     err,
	}
	return c.Status(statusCode).JSON(response)
}

func Success(c *fiber.Ctx, message string, data any) error {
	return SendResponse(c, fiber.StatusOK, data, message, nil)
}

func Created(c *fiber.Ctx, message string, data any) error {
	return SendResponse(c, fiber.StatusCreated, data, message, nil)
}

func BadRequest(c *fiber.Ctx, message string, data any) error {
	return SendResponse(c, fiber.StatusBadRequest, data, message, nil)
}

func Unauthorized(c *fiber.Ctx, message string, data any) error {
	return SendResponse(c, fiber.StatusUnauthorized, data, message, nil)
}

func Forbidden(c *fiber.Ctx, message string, data any) error {
	return SendResponse(c, fiber.StatusForbidden, data, message, nil)
}

func NotFound(c *fiber.Ctx, message string, data any) error {
	return SendResponse(c, fiber.StatusNotFound, data, message, nil)
}

func InternalServerError(c *fiber.Ctx, message string, data any) error {
	return SendResponse(c, fiber.StatusInternalServerError, data, message, nil)
}
