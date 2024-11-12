package response

import "github.com/gofiber/fiber/v2"

// SendResponse is a helper to send JSON responses in Fiber
func SendResponse(c *fiber.Ctx, statusCode int, data interface{}, message string) error {
	response := fiber.Map{
		"statusCode": statusCode,
		"data":       data,
		"message":    message,
	}
	return c.Status(statusCode).JSON(response)
}
