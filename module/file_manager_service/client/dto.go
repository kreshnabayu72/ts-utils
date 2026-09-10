package client

// filePartName is the multipart part that carries the file bytes. One file per request, so the
// name carries no index.
const filePartName = "files"

const maxErrorBodyBytes = 4 << 10

// UploadParams carries the file_manager_data columns the caller owns. The file-derived columns
// (name, path, format, size) are produced by the server, and the tenant and actor travel in
// headers, so neither appears here.
//
// AssetID, UserID and ModifiedByID are strings on the wire for the reason every id in this
// codebase is: a Snowflake decoded into an `any` becomes a float64 and silently rounds. Leaving
// UserID or ModifiedByID empty lets the server default them to the X-User-ID actor.
type UploadParams struct {
	FolderName   string
	Type         string
	Module       string
	Platform     string
	AssetID      string
	UserID       string
	ModifiedByID string
}

// UploadFileResponse is the created file_manager_data row. ID is the Snowflake as a string (see
// UploadParams). PublicPath is the AES-GCM storage token — opaque to callers, and the only value
// GET /file-manager/:token accepts.
type UploadFileResponse struct {
	ID         string  `json:"_id"`
	PublicPath string  `json:"publicPath"`
	FileName   string  `json:"fileName"`
	FileFormat string  `json:"fileFormat"`
	FileSize   string  `json:"fileSize"`
	FileSizeKb float64 `json:"fileSizeKb"`
}
