package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUploadWritesMetadataBeforeFilePart(t *testing.T) {
	var (
		capturedBody   []byte
		capturedHeader http.Header
	)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedBody, _ = io.ReadAll(r.Body)
		capturedHeader = r.Header.Clone()
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "ok",
			"data":    UploadFileResponse{ID: "123", PublicPath: "token", FileName: "a.txt"},
		})
	}))
	defer srv.Close()

	client := NewFileManagerClient(srv.URL, "secret")
	res, err := client.Upload(
		context.Background(),
		"client-1", "user-9",
		UploadParams{FolderName: "attachments", Type: "asset", Platform: "web", AssetID: "42"},
		"a.txt", "text/plain",
		strings.NewReader("hello"),
	)
	require.NoError(t, err)
	require.Equal(t, "123", res.ID)
	require.Equal(t, "token", res.PublicPath)

	require.Equal(t, "secret", capturedHeader.Get("X-Service-Token"))
	require.Equal(t, "client-1", capturedHeader.Get("X-Client-ID"))
	require.Equal(t, "user-9", capturedHeader.Get("X-User-ID"))

	mediaType, params, err := mime.ParseMediaType(capturedHeader.Get("Content-Type"))
	require.NoError(t, err)
	require.Equal(t, "multipart/form-data", mediaType)

	reader := multipart.NewReader(bytes.NewReader(capturedBody), params["boundary"])
	var (
		order    []string
		fields   = map[string]string{}
		fileBody []byte
	)

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		order = append(order, part.FormName())

		if part.FileName() != "" {
			fileBody, err = io.ReadAll(part)
			require.NoError(t, err)
			continue
		}
		value, err := io.ReadAll(part)
		require.NoError(t, err)
		fields[part.FormName()] = string(value)
	}

	require.Equal(t, "42", fields["assetId"])
	require.Equal(t, "attachments", fields["folderName"])
	require.Equal(t, "hello", string(fileBody))

	// The server names the object from folderName, so the file part must arrive last.
	require.Equal(t, filePartName, order[len(order)-1])
	require.NotContains(t, order[:len(order)-1], filePartName)
}

func TestUploadReturnsErrorOnUnexpectedStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, `{"message":"bad request"}`)
	}))
	defer srv.Close()

	client := NewFileManagerClient(srv.URL, "secret")
	_, err := client.Upload(
		context.Background(),
		"client-1", "user-9",
		UploadParams{FolderName: "attachments", Type: "asset", Platform: "web"},
		"a.txt", "text/plain",
		strings.NewReader("hello"),
	)
	require.Error(t, err)
	require.Contains(t, err.Error(), "400")
}
