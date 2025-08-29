package file

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

type FileUtilImpl struct {
	Host string
}

type uploadBody struct {
	DirName    string `json:"dirName"`
	FilePrefix string `json:"filePrefix"`
}

type saveFileResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data,omitempty"`
}

func (f *FileUtilImpl) SaveFile(
	folderName string,
	filePrefix string,
	files []*multipart.FileHeader,
) ([]string, error) {
	route := fmt.Sprintf("%s/api/file/upload", f.Host)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Extra form fields
	_ = writer.WriteField("dirName", folderName)
	_ = writer.WriteField("filePrefix", filePrefix)

	// Attach each file
	for _, fh := range files {
		src, err := fh.Open()
		if err != nil {
			return nil, fmt.Errorf("open file %s: %w", fh.Filename, err)
		}

		part, err := writer.CreateFormFile("files", filePrefix+fh.Filename)
		if err != nil {
			src.Close()
			return nil, fmt.Errorf("create form file: %w", err)
		}

		if _, err := io.Copy(part, src); err != nil {
			src.Close()
			return nil, fmt.Errorf("copy file: %w", err)
		}
		src.Close()
	}

	// Close writer
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("close writer: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest(http.MethodPost, route, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("upload failed (%d): %s", resp.StatusCode, respBody)
	}

	// Read response (assuming backend returns JSON array of URLs)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	var apiResp saveFileResponse
	if err := json.Unmarshal(respBody, &apiResp); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	return apiResp.Data, nil
}
