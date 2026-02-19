package replica

import "time"

type FileManagerData struct {
	ID                              int       `json:"_id"                                column:"file_manager_data.id"                                    sqlType:"BIGINT"`
	Type                            string    `json:"type"                               column:"file_manager_data.type"                                  sqlType:"VARCHAR"`
	FixedAssetGroupID               int       `json:"fixedAssetGroupId"                  column:"file_manager_data.fixed_asset_group_id"                  sqlType:"BIGINT"`
	TransactionID                   string    `json:"transactionId"                      column:"file_manager_data.transaction_id"                        sqlType:"VARCHAR"`
	FileName                        string    `json:"fileName"                           column:"file_manager_data.file_name"                             sqlType:"VARCHAR"`
	FilePath                        string    `json:"filePath"                           column:"file_manager_data.file_path"                             sqlType:"VARCHAR"`
	FileFormat                      string    `json:"fileFormat"                         column:"file_manager_data.file_format"                           sqlType:"VARCHAR"`
	FileSize                        string    `json:"fileSize"                           column:"file_manager_data.file_size"                             sqlType:"VARCHAR"`
	FileSizeKb                      float64   `json:"fileSizeKb"                         column:"file_manager_data.file_size_kb"                          sqlType:"NUMERIC"`
	UserID                          int       `json:"userId"                             column:"file_manager_data.user_id"                               sqlType:"BIGINT"`
	AttachmentID                    int       `json:"attachmentId"                       column:"file_manager_data.attachment_id"                         sqlType:"BIGINT"`
	AssetCustomFieldID              int       `json:"assetCustomFieldId"                 column:"file_manager_data.asset_custom_field_id"                 sqlType:"BIGINT"`
	FixedAssetDisposalTransactionID int       `json:"fixedAssetDisposalTransactionId"    column:"file_manager_data.fixed_asset_disposal_transaction_id"   sqlType:"BIGINT"`
	AssetID                         int       `json:"assetId"                            column:"file_manager_data.asset_id"                              sqlType:"BIGINT"`
	ModifiedByID                    int       `json:"modifiedById"                       column:"file_manager_data.modified_by_id"                        sqlType:"BIGINT"`
	IsDeleted                       bool      `json:"isDeleted"                          column:"file_manager_data.is_deleted"                            sqlType:"BOOLEAN"`
	DeletedAt                       time.Time `json:"deletedAt"                          column:"file_manager_data.deleted_at"                            sqlType:"TIMESTAMPTZ"`
	CreatedAt                       time.Time `json:"createdAt"                          column:"file_manager_data.created_at"                            sqlType:"TIMESTAMPTZ"`
	UpdatedAt                       time.Time `json:"updatedAt"                          column:"file_manager_data.updated_at"                            sqlType:"TIMESTAMPTZ"`
}

func (f *FileManagerData) GetTableName() string {
	return "file_manager_data"
}

func (f *FileManagerData) GetTableNameWithSchema() string {
	return "public.file_manager_data"
}

func (f *FileManagerData) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":                                  "BIGINT",
		"type":                                "VARCHAR",
		"fixed_asset_group_id":                "BIGINT",
		"transaction_id":                      "VARCHAR",
		"file_name":                           "VARCHAR",
		"file_path":                           "VARCHAR",
		"file_format":                         "VARCHAR",
		"file_size":                           "VARCHAR",
		"file_size_kb":                        "NUMERIC",
		"user_id":                             "BIGINT",
		"attachment_id":                       "BIGINT",
		"asset_custom_field_id":               "BIGINT",
		"fixed_asset_disposal_transaction_id": "BIGINT",
		"asset_id":                            "BIGINT",
		"modified_by_id":                      "BIGINT",
		"is_deleted":                          "BOOLEAN",
		"deleted_at":                          "TIMESTAMPTZ",
		"created_at":                          "TIMESTAMPTZ",
		"updated_at":                          "TIMESTAMPTZ",
	}
}

func (f *FileManagerData) GetColumns() []string {
	return []string{
		"id",
		"type",
		"fixed_asset_group_id",
		"transaction_id",
		"file_name",
		"file_path",
		"file_format",
		"file_size",
		"file_size_kb",
		"user_id",
		"attachment_id",
		"asset_custom_field_id",
		"fixed_asset_disposal_transaction_id",
		"asset_id",
		"modified_by_id",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}
