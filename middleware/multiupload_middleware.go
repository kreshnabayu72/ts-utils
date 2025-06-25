package middleware

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/gofiber/fiber/v2"
	_ "golang.org/x/image/webp"
)

// MultiUploadMiddleware handles multi uploads and saves them to disk.
func MultiUploadMiddleware(folderName string, attributes []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Get the uploaded file, FormFile attribute
		for _, attribute := range attributes {
			fileHeader, err := ctx.FormFile(attribute)
			if errors.Is(err, fasthttp.ErrMissingFile) {
				continue
			} else if err != nil {
				return err
			}

			file, err := fileHeader.Open()
			if err != nil {
				file.Close()

				return fiber.NewError(fiber.StatusBadRequest, "Error retrieving the file: "+err.Error())
			}

			companyCode := ctx.Locals(ContextKey("companyCode")).(string)

			absPath, _ := filepath.Abs(
				fmt.Sprintf("../../tsv2-files/%s/%s", companyCode, folderName),
			)
			// Save the resized image to disk
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
				file.Close()

				return fiber.NewError(
					fiber.StatusInternalServerError,
					"Error creating directory: "+err.Error(),
				)
			}

			err = ctx.SaveFile(fileHeader, outputFilePath)
			if err != nil {
				file.Close()

				return fiber.NewError(
					fiber.StatusInternalServerError,
					"Error saving resized image: "+err.Error(),
				)
			}

			key := ContextKey(attribute)
			path := fmt.Sprintf("/%s/%s/%s", companyCode, folderName, filename)

			ctx.Locals(key, path)

			file.Close()
		}

		return ctx.Next()
	}
}
