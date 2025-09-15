package file

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type FileUtil interface {
	SaveFile(
		ctx *fiber.Ctx,
		folderName string,
		filePrefix string,
		files []*multipart.FileHeader,
	) ([]string, error)
}
