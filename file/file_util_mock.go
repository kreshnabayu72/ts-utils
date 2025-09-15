package file

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

type MockFileUtil struct {
	mock.Mock
}

func (m *MockFileUtil) SaveFile(
	ctx *fiber.Ctx,
	folderName string,
	filePrefix string,
	files []*multipart.FileHeader,
) ([]string, error) {
	args := m.Called(ctx, folderName, filePrefix, files)
	return args.Get(0).([]string), args.Error(1)
}
