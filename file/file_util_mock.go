package file

import (
	"mime/multipart"

	"github.com/stretchr/testify/mock"
)

type MockFileUtil struct {
	mock.Mock
}

func (m *MockFileUtil) SaveFile(
	folderName string,
	filePrefix string,
	files []*multipart.FileHeader,
) ([]string, error) {
	args := m.Called(folderName, filePrefix, files)
	return args.Get(0).([]string), args.Error(1)
}
