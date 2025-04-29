package functions

import (
	"errors"
	"fmt"
	"gitea.qwertysystem.net/BETS/ts-utils/middleware"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// SaveFile saves the provided file from multipart/form-data and store them in the disk.
//
// This function always treat the request attribute field of the file as optional,
// if consumer doesn't provide any file, it won't return any error.
func SaveFile(ctx *fiber.Ctx, folderName string, attribute string) (string, error) {
	fileHeader, err := ctx.FormFile(attribute)

	// If payload doesn't provide any file, skip the subsequent execution
	if errors.Is(err, fasthttp.ErrMissingFile) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", fiber.NewError(fiber.StatusBadRequest, "Error retrieving the file: "+err.Error())
	}

	defer file.Close()

	companyCode := ctx.Locals(middleware.ContextKey("companyCode")).(string)

	absPath, _ := filepath.Abs(
		fmt.Sprintf("../../../../tsv2-files/%s/%s", companyCode, folderName),
	)

	filename := fmt.Sprintf(
		"%s.%s",
		strconv.FormatInt(time.Now().Unix(), 10)+"-"+attribute,
		fileHeader.Filename,
	)
	outputFilePath := filepath.Join(
		absPath,
		filename,
	)

	// Create the directory if it doesn't exist
	if err = os.MkdirAll(absPath, os.ModePerm); err != nil {
		return "", fiber.NewError(
			fiber.StatusInternalServerError,
			"Error creating directory: "+err.Error(),
		)
	}

	err = ctx.SaveFile(fileHeader, outputFilePath)
	if err != nil {
		return "", fiber.NewError(
			fiber.StatusInternalServerError,
			"Error saving resized image: "+err.Error(),
		)
	}

	path := fmt.Sprintf("/%s/%s/%s", companyCode, folderName, filename)

	return path, nil
}
