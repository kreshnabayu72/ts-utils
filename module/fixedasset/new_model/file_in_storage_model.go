package new_model

import "time"

type FileInStorage struct {
	ID                   int       `json:"_id" column:"id"`
	AttachmentID         *int      `json:"attachmentId" column:"attachment_id"`
	AssetID              int       `json:"assetId" column:"asset_id"`
	ModifiedByID         int       `json:"modifiedById" column:"modified_by_id"`
	TransactionID        *string   `json:"transactionId" column:"transaction_id"`
	Type                 FileType  `json:"type" column:"type"`
	FileName             string    `json:"fileName" column:"file_name"`
	FilePath             string    `json:"filePath" column:"file_path"`
	FileFormat           string    `json:"fileFormat" column:"file_format"`
	FileSize             string    `json:"fileSize" column:"file_size"`
	FileSizeKB           float64   `json:"fileSizeKb" column:"file_size_kb"`
	IsDeleted            bool      `json:"isDeleted" column:"is_deleted"`
	DeleteFromDatabaseAt time.Time `json:"deleteFromDatabaseAt" column:"delete_from_database_at"`
	CreatedAt            time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt            time.Time `json:"updatedAt" column:"updated_at"`
}
