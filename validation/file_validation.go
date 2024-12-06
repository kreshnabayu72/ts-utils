package validation

import (
	"fmt"
	"mime/multipart"
	"strings"
)

func ValidateFile(file *multipart.FileHeader, allowedTypes []string, maxSize int64) error {
	// Validasi ukuran file
	if file.Size > maxSize {
		return fmt.Errorf("file size exceeds %d bytes", maxSize)
	}

	// Validasi tipe file
	contentType := file.Header.Get("Content-Type")
	for _, allowed := range allowedTypes {
		if strings.EqualFold(contentType, allowed) {
			return nil
		}
	}

	return fmt.Errorf("invalid file type: %s", contentType)
}
