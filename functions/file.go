package functions

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/middleware"

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

// SaveFileWithName saves the uploaded file (if any) into
// ../../../../tsv2-files/{companyCode}/{folderName}/{fileName}.
// If the provided fileName has no extension, the original file’s
// extension is appended. If no file is uploaded under `attribute`,
// it returns ("", nil).
func SaveFileWithName(
	ctx *fiber.Ctx,
	folderName string,
	attribute string,
	fileName string,
) (string, error) {
	// 1. Try to get the uploaded file
	fileHeader, err := ctx.FormFile(attribute)
	if errors.Is(err, fasthttp.ErrMissingFile) {
		// no file provided → nothing to do
		return "", nil
	}
	if err != nil {
		return "", err
	}

	// 2. Ensure fileName has an extension
	origExt := filepath.Ext(fileHeader.Filename)
	if filepath.Ext(fileName) == "" {
		fileName += origExt
	}

	// 3. Build absolute path
	companyCode := ctx.Locals(middleware.ContextKey("companyCode")).(string)
	baseDir := fmt.Sprintf("../../../../tsv2-files/%s/%s", companyCode, folderName)
	absPath, err := filepath.Abs(baseDir)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError,
			"Error resolving storage path: "+err.Error())
	}

	// 4. Ensure directory exists
	if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError,
			"Error creating directory: "+err.Error())
	}

	// 5. Save the file
	dst := filepath.Join(absPath, fileName)
	if err := ctx.SaveFile(fileHeader, dst); err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError,
			"Error saving file: "+err.Error())
	}

	// 6. Return the public path
	publicPath := fmt.Sprintf("/%s/%s/%s", companyCode, folderName, fileName)
	return publicPath, nil

}

func DeleteFile(filePath string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	os.Remove(filepath.Join(currentDir, "..", "..", "..", "tsv2-files", filePath))

	return nil
}

func DeleteSplittedFile(fileMediumPath string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	os.Remove(filepath.Join(currentDir, "..", "..", "..", "tsv2-files", fileMediumPath))
	big := strings.ReplaceAll(fileMediumPath, "medium", "big")
	os.Remove(filepath.Join(currentDir, "..", "..", "..", "tsv2-files", big))
	small := strings.ReplaceAll(fileMediumPath, "medium", "small")
	os.Remove(filepath.Join(currentDir, "..", "..", "..", "tsv2-files", small))

	return nil
}
