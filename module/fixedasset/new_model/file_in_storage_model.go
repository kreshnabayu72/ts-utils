package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type FileInStorage struct {
	ID                   int       `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AttachmentID         *int      `json:"attachmentId" column:"attachment_id"` // not used in SQL, kept for app logic
	AssetID              int       `json:"assetId" column:"asset_id" sql:"BIGINT"`
	ModifiedByID         int       `json:"modifiedById" column:"modified_by_id" sql:"BIGINT"`
	TransactionID        *string   `json:"transactionId" column:"transaction_id" sql:"VARCHAR"`
	Type                 FileType  `json:"type" column:"type" sql:"file_type_enum"`
	FileName             string    `json:"fileName" column:"file_name" sql:"VARCHAR"`
	FilePath             string    `json:"filePath" column:"file_path" sql:"VARCHAR"`
	FileFormat           string    `json:"fileFormat" column:"file_format" sql:"VARCHAR"`
	FileSize             string    `json:"fileSize" column:"file_size" sql:"VARCHAR"`
	FileSizeKB           float64   `json:"fileSizeKb" column:"file_size_kb" sql:"FLOAT"`
	IsDeleted            bool      `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN"`
	DeleteFromDatabaseAt time.Time `json:"deleteFromDatabaseAt" column:"deleted_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt            time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt            time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (f *FileInStorage) TableName() string {
	return "file_in_storage"
}

func (f *FileInStorage) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(f)

	builder.AddForeignKey(
		"fk_file_in_storage_asset_id",
		"asset_id",
		"assets",
		"id",
		"CASCADE",
	)

	builder.AddForeignKey(
		"fk_file_in_storage_modified_by_id",
		"modified_by_id",
		"users",
		"id",
		"CASCADE",
	)
}
