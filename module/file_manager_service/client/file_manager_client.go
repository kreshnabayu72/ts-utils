// Package client holds the file-manager HTTP contract and the outbound client that speaks it,
// so every producer of an upload shares one implementation instead of copying it.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
)

// FileManagerClient stores files through file-manager-service. It authenticates as this service
// with X-Service-Token; the tenant and the acting user travel in X-Client-ID and X-User-ID, so a
// service token alone can never write into another tenant's schema.
type FileManagerClient struct {
	baseURL      string
	serviceToken string
	http         *http.Client
}

// NewFileManagerClient builds a client against file-manager-service's internal base URL. The
// timeout is far longer than a request/response call needs because the request body is the whole
// file.
func NewFileManagerClient(baseURL, serviceToken string) *FileManagerClient {
	return &FileManagerClient{
		baseURL:      baseURL,
		serviceToken: serviceToken,
		http:         &http.Client{Timeout: 30 * time.Second},
	}
}

// Upload stores one file and creates its file_manager_data row in the caller's tenant
// (POST /internal/file-manager/upload). filename becomes the row's file_name and the base of the
// object key; contentType is stored on the object and echoed back when the file is read.
//
// Metadata parts are written before the file part by construction: a streamed reader yields the
// file part last, and the server needs folderName before it can name the object.
func (c *FileManagerClient) Upload(
	ctx context.Context,
	clientID, userID string,
	params UploadParams,
	filename, contentType string,
	file io.Reader,
) (*UploadFileResponse, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	fields := []struct{ name, value string }{
		{"folderName", params.FolderName},
		{"type", params.Type},
		{"module", params.Module},
		{"platform", params.Platform},
		{"assetId", params.AssetID},
		{"userId", params.UserID},
		{"modifiedById", params.ModifiedByID},
	}
	for _, field := range fields {
		if field.value == "" {
			continue
		}
		if err := writer.WriteField(field.name, field.value); err != nil {
			return nil, fmt.Errorf("filemanager.Upload: write field %s: %w", field.name, err)
		}
	}

	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", mime.FormatMediaType("form-data", map[string]string{
		"name":     filePartName,
		"filename": filename,
	}))
	header.Set("Content-Type", contentType)
	part, err := writer.CreatePart(header)
	if err != nil {
		return nil, fmt.Errorf("filemanager.Upload: create file part: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("filemanager.Upload: copy file: %w", err)
	}
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("filemanager.Upload: close writer: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.baseURL+constant.FILE_MANAGER_INTERNAL_UPLOAD_PATH,
		&body,
	)
	if err != nil {
		return nil, fmt.Errorf("filemanager.Upload: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set(constant.HEADER_SERVICE_TOKEN, c.serviceToken)
	req.Header.Set(constant.HEADER_CLIENT_ID, clientID)
	req.Header.Set(constant.HEADER_USER_ID, userID)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("filemanager.Upload: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		detail, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrorBodyBytes))
		return nil, fmt.Errorf("filemanager.Upload: unexpected status %d: %s", resp.StatusCode, detail)
	}

	var envelope struct {
		Data UploadFileResponse `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		return nil, fmt.Errorf("filemanager.Upload: decode response: %w", err)
	}
	return &envelope.Data, nil
}
