package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// Asset
// - DepreciationGroupID:
// For future reference, this is used for the basic version.
// Basic version bypass the creation of depreciation group through asset name, so we set it immediately to the asset.
type Asset struct {
	ID                     int               `json:"_id"                      column:"assets.id"                         sql:"BIGINT PRIMARY KEY"`
	AssetIDFull            string            `json:"assetIdFull"              column:"assets.asset_id_full"              sql:"VARCHAR UNIQUE"`
	AssetID                string            `json:"assetId"                  column:"assets.asset_id"                   sql:"VARCHAR UNIQUE"`
	NameID                 int               `json:"nameId"                   column:"assets.name_id"                    sql:"BIGINT"`
	NameWithSequence       string            `json:"nameWithSequence"         column:"assets.name_with_sequence"         sql:"VARCHAR"`
	SequenceNumber         int               `json:"sequenceNumber"           column:"assets.sequence_number"            sql:"BIGINT"`
	AliasName              *string           `json:"aliasName"                column:"assets.alias_name"                 sql:"VARCHAR"`
	AliasCode              *string           `json:"aliasCode"                column:"assets.alias_code"                 sql:"VARCHAR"`
	QrID                   *int              `json:"qrId"                     column:"assets.qr_id"                      sql:"BIGINT"`
	RfidID                 *int              `json:"rfidId"                   column:"assets.rfid_id"                    sql:"BIGINT"`
	Status                 AssetStatus       `json:"status"                   column:"assets.status"                     sql:"asset_status_enum NOT NULL"`
	FirstImage             *string           `json:"firstImage"               column:"assets.first_image"                sql:"VARCHAR"`
	SecondImage            *string           `json:"secondImage"              column:"assets.second_image"               sql:"VARCHAR"`
	SetDefault             AssetDefaultImage `json:"setDefault"               column:"assets.set_default"                sql:"asset_default_image_enum"`
	PreviousStatus         *string           `json:"previousStatus"           column:"assets.previous_status"            sql:"asset_status_enum"`
	TagType                TagType           `json:"tagType"                  column:"assets.tag_type"                   special:"generated"`
	GroupID                int               `json:"groupId"                  column:"assets.group_id"                   sql:"BIGINT"`
	ModelID                int               `json:"modelId"                  column:"assets.model_id"                   sql:"BIGINT"`
	AssignedToID           *int              `json:"assignedToId"             column:"assets.assigned_to_id"             sql:"BIGINT"`
	BorrowerID             *int              `json:"borrowerId"               column:"assets.borrower_id"                sql:"BIGINT"`
	LastModifiedID         *int              `json:"lastModifiedId"           column:"assets.last_modified_id"           sql:"BIGINT"`
	ParentID               *int              `json:"parentId"                 column:"assets.parent_id"                  sql:"BIGINT"`
	DepreciationGroupID    *int              `json:"depreciationGroupId"      column:"assets.depreciation_group_id"      sql:"BIGINT"`
	PurchaseDate           *time.Time        `json:"purchaseDate"             column:"assets.purchase_date"              sql:"TIMESTAMPTZ"`
	PurchaseCost           *float64          `json:"purchaseCost"             column:"assets.purchase_cost"              sql:"FLOAT"`
	PurchaseCurrency       *string           `json:"purchaseCurrency"         column:"assets.purchase_currency"          sql:"VARCHAR"`
	PurchaseWarrantyYears  *int              `json:"purchaseWarrantyDuration" column:"assets.purchase_warranty_duration" sql:"INTEGER"`
	PurchaseCurrencySymbol *string           `json:"purchaseCurrencySymbol"   column:"assets.purchase_currency_symbol"   sql:"VARCHAR"`
	IsDeleted              bool              `json:"isDeleted"                column:"assets.is_deleted"                 sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	LastScannedAt          *time.Time        `json:"lastScannedAt"            column:"assets.last_scanned_at"            sql:"TIMESTAMPTZ DEFAULT NOW()"`
	CreatedAt              time.Time         `json:"createdAt"                column:"assets.created_at"                 sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt              time.Time         `json:"updatedAt"                column:"assets.updated_at"                 sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Asset) TableName() string {
	return "assets"
}

func (m *Asset) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey("fk_assets_name", "name_id", "asset_names", "id", "RESTRICT")
	builder.AddForeignKey("fk_assets_model", "model_id", "models", "id", "RESTRICT")
	builder.AddForeignKey("fk_assets_group", "group_id", "groups", "id", "RESTRICT")
	builder.AddForeignKey("fk_assets_qr", "qr_id", "qrs", "id", "SET NULL")
	builder.AddForeignKey("fk_assets_rfid", "rfid_id", "rfids", "id", "SET NULL")
	builder.AddForeignKey("fk_assets_parent", "parent_id", "assets", "id", "SET NULL")
	builder.AddForeignKey("fk_assets_last_modified", "last_modified_id", "users", "id", "SET NULL")
	builder.AddForeignKey("fk_assets_assigned_to", "assigned_to_id", "users", "id", "SET NULL")
	builder.AddForeignKey("fk_assets_borrower", "borrower_id", "users", "id", "SET NULL")

	builder.AddIndex("idx_assets_asset_id", "asset_id", false, "")
	builder.AddIndex("idx_assets_name_id", "name_id", false, "")
	builder.AddIndex("idx_assets_model_id", "model_id", false, "")
	builder.AddIndex("idx_assets_group_id", "group_id", false, "")
	builder.AddIndex("idx_assets_parent_id", "parent_id", false, "")
	builder.AddIndex("idx_assets_status", "status", false, "")
	builder.AddIndex("idx_assets_is_deleted", "is_deleted", false, "")
	builder.AddIndex("idx_assets_sequence_number", "sequence_number", false, "")
}
