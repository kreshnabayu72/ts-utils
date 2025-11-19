package model

import (
	"time"
)

type FileInStorage struct {
	ID                   int       `json:"_id"                  column:"file_in_storage.id"`
	AttachmentID         *int      `json:"attachmentId"         column:"file_in_storage.attachment_id"` // not used in SQL, kept for app logic - UPDATE(kreshna): used for delete attachment API so it can delete file_in_storage & update storage kb
	AssetID              int       `json:"assetId"              column:"file_in_storage.asset_id"`
	ModifiedByID         int       `json:"modifiedById"         column:"file_in_storage.modified_by_id"`
	TransactionID        *string   `json:"transactionId"        column:"file_in_storage.transaction_id"`
	Type                 FileType  `json:"type"                 column:"file_in_storage.type"`
	FileName             string    `json:"fileName"             column:"file_in_storage.file_name"`
	FilePath             string    `json:"filePath"             column:"file_in_storage.file_path"`
	FileFormat           string    `json:"fileFormat"           column:"file_in_storage.file_format"`
	FileSize             string    `json:"fileSize"             column:"file_in_storage.file_size"`
	FileSizeKB           float64   `json:"fileSizeKb"           column:"file_in_storage.file_size_kb"`
	IsDeleted            bool      `json:"isDeleted"            column:"file_in_storage.is_deleted"`
	DeleteFromDatabaseAt time.Time `json:"deleteFromDatabaseAt" column:"file_in_storage.deleted_at"`
	CreatedAt            time.Time `json:"createdAt"            column:"file_in_storage.created_at"`
	UpdatedAt            time.Time `json:"updatedAt"            column:"file_in_storage.updated_at"`
}
