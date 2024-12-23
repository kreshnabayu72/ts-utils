package validation

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

type MockValidator struct {
	mock.Mock
}

func (m *MockValidator) ValidateFile(
	file *multipart.FileHeader,
	allowedTypes []string,
	maxSize int64,
) error {
	args := m.Called(file, allowedTypes, maxSize)
	return args.Error(0)
}

func (m *MockValidator) ValidateStruct(c *fiber.Ctx, payload interface{}) error {
	args := m.Called(c, payload)
	return args.Error(0)
}
