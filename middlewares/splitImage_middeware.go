package middlewares

import (
	"context"
	"fmt"
	"image"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/betsv2/tsv2-go-utils/types"

	"github.com/disintegration/imaging"
)

// uploadHandler handles image upload and resizing
// splitImageMiddleware handles image uploads, resizes the images, and saves them to disk.
func SplitImageMiddleware(next http.HandlerFunc, folderName string, attribute string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Limit the maximum size of an uploaded file (optional)
		err := r.ParseMultipartForm(10 << 20) // Max 10MB
		if err != nil {
			http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Get the uploaded file, FormFile attribute
		file, _, err := r.FormFile(attribute)
		if err != nil {
			http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Decode the image
		img, _, err := image.Decode(file)
		if err != nil {
			http.Error(w, "Error decoding image: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Define sizes for resizing the image
		sizes := []struct {
			name          string
			width, height int
		}{
			{"small", 300, 300},
			{"medium", 600, 600},
			{"large", 1200, 1200},
		}

		// Start with the original context
		ctx := r.Context()

		// Process and save resized images
		for _, size := range sizes {
			// Resize the image to the specified dimensions
			resizedImg := imaging.Resize(img, size.width, size.height, imaging.Lanczos)

			companyCode := ctx.Value(types.ContextKey("companyCode"))

			absPath, _ := filepath.Abs(fmt.Sprintf("../tsv2-files/%s/%s", companyCode, folderName))
			// Save the resized image to disk
			outputFilePath := filepath.Join(absPath, fmt.Sprintf("%s.jpg", strconv.FormatInt(time.Now().Unix(), 10)+"-"+size.name))

			// Create the directory if it doesn't exist
			if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
				http.Error(w, "Error creating directory: "+err.Error(), http.StatusInternalServerError)
				return
			}

			err := imaging.Save(resizedImg, outputFilePath)
			if err != nil {
				http.Error(w, "Error saving resized image: "+err.Error(), http.StatusInternalServerError)
				return
			}

			key := types.ContextKey(size.name)

			ctx = context.WithValue(ctx, key, outputFilePath)

		}

		// Response after successfully processing the images
		w.Write([]byte("Image uploaded and resized successfully"))

		r = r.WithContext(ctx)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
