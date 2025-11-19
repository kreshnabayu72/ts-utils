package service

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

// MockFileService implements FileServicer
//
// Used to mock the FileService implementation.
type MockFileService struct {
	mock.Mock
}

func (m *MockFileService) Init() {
	// Pass
}

func (m *MockFileService) SaveFile(
	ctx *fiber.Ctx,
	folderName string,
	filePrefix string,
	files []*multipart.FileHeader,
) ([]string, error) {
	args := m.Called(ctx, folderName, filePrefix, files)
	return args.Get(0).([]string), args.Error(1)
}

func (m *MockFileService) Close() {
	// Pass
}
