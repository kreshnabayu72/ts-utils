package functions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
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

// Define a struct matching the response
type UploadResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

// SaveFileWithName saves the uploaded file (if any) into
// ../../../../tsv2-files/{companyCode}/{folderName}/{fileName}.
// If the provided fileName has no extension, the original file’s
// extension is appended. If no file is uploaded under `attribute`,
func SaveFileWithName(
	ctx *fiber.Ctx,
	folderName string,
	attribute string,
	fileName string,
) (string, string, string, int64, string, error) {
	// 1. Try to get the uploaded file
	fileHeader, err := ctx.FormFile(attribute)
	if errors.Is(err, fasthttp.ErrMissingFile) {
		// no file provided → nothing to do
		return "", "", "", 0, "", nil
	}
	if err != nil {
		return "", "", "", 0, "", err
	}

	// 2. Ensure fileName has an extension
	origExt := filepath.Ext(fileHeader.Filename)
	if filepath.Ext(fileName) == "" {
		fileName += origExt
	}

	// 3. Build absolute path
	companyCode := ctx.Locals(middleware.ContextKey("companyCode")).(string)
	baseDir := fmt.Sprintf("../../tsv2-files/%s/%s", companyCode, folderName)
	absPath, err := filepath.Abs(baseDir)
	if err != nil {
		return "", "", "", 0, "", fiber.NewError(fiber.StatusInternalServerError,
			"Error resolving storage path: "+err.Error())
	}

	// 4. Ensure directory exists
	if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
		return "", "", "", 0, "", fiber.NewError(fiber.StatusInternalServerError,
			"Error creating directory: "+err.Error())
	}

	// 5. Save the file

	// Open file from memory
	src, err := fileHeader.Open()
	if err != nil {
		return "", "", "", 0, "", fiber.NewError(fiber.StatusInternalServerError, "Cannot open uploaded file: "+err.Error())
	}
	defer src.Close()

	// Create multipart body
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Create form file part named "files"
	part, err := writer.CreateFormFile("files", fileHeader.Filename)
	if err != nil {
		return "", "", "", 0, "", fiber.NewError(fiber.StatusInternalServerError, "Cannot create form file part: "+err.Error())
	}

	//other form part
	writer.WriteField("dirName", folderName)
	writer.WriteField("filePrefix", folderName)

	// Copy the uploaded file content into the request body
	if _, err := io.Copy(part, src); err != nil {
		return "", "", "", 0, "", fiber.NewError(fiber.StatusInternalServerError, "Cannot copy file: "+err.Error())
	}

	writer.Close()

	fileApi := os.Getenv("FILE_MANAGER_API")
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/api/file/upload", fileApi), &b)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	if err != nil {
		return "", "", "", 0, "", fiber.NewError(fiber.StatusInternalServerError,
			"Error saving file: "+err.Error())

	}

	// Forward authorization header if present
	if auth := ctx.Get("Authorization"); auth != "" {
		req.Header.Set("Authorization", auth)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", "", 0, "", fiber.NewError(fiber.StatusInternalServerError, "Upload failed: "+err.Error())
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", "", "", 0, "", fiber.NewError(resp.StatusCode, "Upload failed: "+string(respBody))
	}

	var uploadResp UploadResponse
	if err := json.Unmarshal(respBody, &uploadResp); err != nil {
		return "", "", "", 0, "", fiber.NewError(fiber.StatusInternalServerError, "Failed to parse upload response: "+err.Error())
	}

	//this batch is replaced
	dst := filepath.Join(absPath, fileName)
	if err := ctx.SaveFile(fileHeader, dst); err != nil {
		return "", "", "", 0, "", fiber.NewError(fiber.StatusInternalServerError,
			"Error saving file: "+err.Error())
	}

	var filetype string
	if fileType, err := middleware.GetFileType(fileHeader.Header.Get("Content-Type")); err == nil {
		filetype = fileType
	}

	// 6. Return the public path
	publicPath := fmt.Sprintf("/%s/%s/%s", companyCode, folderName, fileName)

	if len(uploadResp.Data) > 0 {
		publicPath = uploadResp.Data[0]
	}

	return publicPath, fileName, filetype, fileHeader.Size, middleware.GetFileSizeString(uint64(fileHeader.Size)), nil

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
