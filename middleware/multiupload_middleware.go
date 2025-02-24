package middleware

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// uploadHandler handles image upload and resizing
// multiImageMiddleware handles multi image uploads and saves them to disk.
func MultiImageMiddleware(folderName string, attributes []string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// Get the uploaded file, FormFile attribute
		for _, attribute := range attributes {
			fileHeader, err := ctx.FormFile(attribute)
			if err != nil {
				for _, attributeErr := range attributes {
					ctx.Locals(ContextKey(attributeErr), "")
				}
				return ctx.Next()
			}

			file, err := fileHeader.Open()
			if err != nil {
				return fiber.NewError(fiber.StatusBadRequest, "Error retrieving the file: "+err.Error())
			}

			defer file.Close()

			companyCode := ctx.Locals(ContextKey("companyCode")).(string)

			absPath, _ := filepath.Abs(
				fmt.Sprintf("../../../../tsv2-files/%s/%s", companyCode, folderName),
			)
			// Save the resized image to disk
			filename := fmt.Sprintf(
				"%s.jpg",
				strconv.FormatInt(time.Now().Unix(), 10)+"-"+attribute,
			)
			outputFilePath := filepath.Join(
				absPath,
				filename,
			)

			// Create the directory if it doesn't exist
			if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
				return fiber.NewError(
					fiber.StatusInternalServerError,
					"Error creating directory: "+err.Error(),
				)
			}

			err = ctx.SaveFile(fileHeader, outputFilePath)
			if err != nil {
				return fiber.NewError(
					fiber.StatusInternalServerError,
					"Error saving resized image: "+err.Error(),
				)
			}

			key := ContextKey(attribute)
			path := fmt.Sprintf("/%s/%s/%s", companyCode, folderName, filename)

			ctx.Locals(key, path)

		}

		return ctx.Next()
	}
}
