// Package model defines the file-manager-service database row types.
package model

import "time"

type FileManagerData struct {
	ID                              int64      `json:"id"                              column:"file_manager_data.id"`
	Type                            string     `json:"type"                            column:"file_manager_data.type"`
	FixedAssetGroupID               *int64     `json:"fixedAssetGroupId"               column:"file_manager_data.fixed_asset_group_id"`
	TransactionID                   *string    `json:"transactionId"                   column:"file_manager_data.transaction_id"`
	FileName                        string     `json:"fileName"                        column:"file_manager_data.file_name"`
	FilePath                        string     `json:"filePath"                        column:"file_manager_data.file_path"`
	FileFormat                      string     `json:"fileFormat"                      column:"file_manager_data.file_format"`
	FileSize                        string     `json:"fileSize"                        column:"file_manager_data.file_size"`
	FileSizeKb                      float64    `json:"fileSizeKb"                      column:"file_manager_data.file_size_kb"`
	UserID                          *int64     `json:"userId"                          column:"file_manager_data.user_id"`
	AttachmentID                    *int64     `json:"attachmentId"                    column:"file_manager_data.attachment_id"`
	AssetCustomFieldID              *int64     `json:"assetCustomFieldId"              column:"file_manager_data.asset_custom_field_id"`
	FixedAssetDisposalTransactionID *int64     `json:"fixedAssetDisposalTransactionId" column:"file_manager_data.fixed_asset_disposal_transaction_id"`
	AssetID                         *int64     `json:"assetId"                         column:"file_manager_data.asset_id"`
	ModifiedByID                    *int64     `json:"modifiedById"                    column:"file_manager_data.modified_by_id"`
	IsDeleted                       bool       `json:"isDeleted"                       column:"file_manager_data.is_deleted"`
	DeletedAt                       *time.Time `json:"deletedAt"                       column:"file_manager_data.deleted_at"`
	CreatedAt                       time.Time  `json:"createdAt"                       column:"file_manager_data.created_at"`
	UpdatedAt                       time.Time  `json:"updatedAt"                       column:"file_manager_data.updated_at"`
}

