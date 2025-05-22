package middleware

import (
	"fmt"
	"image"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/gofiber/fiber/v2"
)

func GetFileType(headerValue string) (string, error) {
	// mime.ParseMediaType will strip any parameters for you
	mediaType, _, err := mime.ParseMediaType(headerValue)
	if err != nil {
		return "", err
	}
	parts := strings.Split(mediaType, "/")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid media type: %q", mediaType)
	}
	return strings.ToUpper(parts[1]), nil
}

// converts a size in bytes into a string like "24 MB", "400 KB", "1 GB", etc.
func GetFileSizeString(bytes uint64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	f := float64(bytes)
	i := 0
	for f >= 1024 && i < len(units)-1 {
		f /= 1024
		i++
	}
	// "%.0f" rounds to nearest integer; drop decimals entirely
	return fmt.Sprintf("%.0f %s", f, units[i])
}

// uploadHandler handles image upload and resizing
// splitImageMiddleware handles image uploads, resizes the images, and saves them to disk.
func SplitImageMiddleware(folderName string, attribute string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Define sizes for resizing the image
		sizes := []struct {
			name          string
			width, height int
		}{
			{"small", 300, 300},
			{"medium", 600, 600},
			{"large", 1200, 1200},
		}

		// Get the uploaded file, FormFile attribute
		fileHeader, err := ctx.FormFile(attribute)
		if err != nil {
			for _, size := range sizes {
				ctx.Locals(ContextKey(size.name), "")
			}
			return ctx.Next()
		}

		fmt.Println("Uploaded file size:", fileHeader.Size)
		fmt.Println("Uploaded file type:", fileHeader.Header.Get("Content-Type"))

		if fileType, err := GetFileType(fileHeader.Header.Get("Content-Type")); err == nil {
			fmt.Println("Uploaded file type:", fileType) // prints "PNG"
			ctx.Locals(FileExt, fileType)
		} else {
			// handle parse error
			fmt.Println("Could not parse content type:", err)
		}

		ctx.Locals(FileSize, fileHeader.Size)
		ctx.Locals(FileSizeString, GetFileSizeString(uint64(fileHeader.Size)))

		file, err := fileHeader.Open()
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Error retrieving the file: "+err.Error())
		}

		defer file.Close()

		// Decode the image
		img, _, err := image.Decode(file)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Error decoding image: "+err.Error())
		}

		// Process and save resized images
		fileNamePrefix := strconv.FormatInt(time.Now().Unix(), 10)
		for _, size := range sizes {
			// Resize the image to the specified dimensions
			resizedImg := imaging.Resize(img, size.width, size.height, imaging.Lanczos)

			companyCode := ctx.Locals(ContextKey("companyCode")).(string)

			absPath, _ := filepath.Abs(
				fmt.Sprintf("../../tsv2-files/%s/%s", companyCode, folderName),
			)

			// Save the resized image to disk
			filename := fmt.Sprintf(
				"%s.jpg",
				fileNamePrefix+"-"+size.name,
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

			err := imaging.Save(resizedImg, outputFilePath)
			if err != nil {
				return fiber.NewError(
					fiber.StatusInternalServerError,
					"Error saving resized image: "+err.Error(),
				)
			}

			key := ContextKey(size.name)
			path := fmt.Sprintf("/%s/%s/%s", companyCode, folderName, filename)

			ctx.Locals(key, path)
		}

		ctx.Locals(FileNameWithoutExt, fileNamePrefix)

		return ctx.Next()
	}
}
