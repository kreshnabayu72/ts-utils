package file

import "mime/multipart"

type FileUtil interface {
	SaveFile(
		folderName string,
		filePrefix string,
		files []*multipart.FileHeader,
	) ([]string, error)
}
