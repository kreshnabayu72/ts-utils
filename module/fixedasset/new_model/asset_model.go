package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Asset struct {
	ID               int               `json:"_id"                column:"id"                 sql:"BIGINT PRIMARY KEY"`
	AssetIDFull      string            `json:"assetIdFull"        column:"asset_id_full"      sql:"VARCHAR UNIQUE"`
	AssetID          string            `json:"assetId"            column:"asset_id"           sql:"VARCHAR UNIQUE"`
	NameID           int               `json:"nameId"             column:"name_id"            sql:"BIGINT"`
	NameWithSequence string            `json:"nameWithSequence"   column:"name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int               `json:"sequenceNumber"     column:"sequence_number"    sql:"BIGINT"`
	AliasName        *string           `json:"aliasName"          column:"alias_name"         sql:"VARCHAR"`
	AliasCode        *string           `json:"aliasCode"          column:"alias_code"         sql:"VARCHAR"`
	QrID             *int              `json:"qrId"               column:"qr_id"              sql:"BIGINT"`
	RfidID           *int              `json:"rfidId"             column:"rfid_id"            sql:"BIGINT"`
	Status           AssetStatus       `json:"status"             column:"status"             sql:"asset_status_enum NOT NULL"`
	FirstImage       *string           `json:"firstImage"         column:"first_image"        sql:"VARCHAR"`
	SecondImage      *string           `json:"secondImage"        column:"second_image"       sql:"VARCHAR"`
	SetDefault       AssetDefaultImage `json:"setDefault"         column:"set_default"        sql:"asset_default_image_enum"`
	PreviousStatus   *string           `json:"previousStatus"     column:"previous_status"    sql:"asset_status_enum NOT NULL"`

	// ─────── References ───────
	GroupID        int  `json:"groupId"        column:"group_id" sql:"BIGINT"`
	ModelID        int  `json:"modelId"        column:"model_id" sql:"BIGINT"`
	AssignedToID   *int `json:"assignedToId"   column:"assigned_to_id" sql:"BIGINT"`
	BorrowerID     *int `json:"borrowerId"     column:"borrower_id" sql:"BIGINT"`
	LastModifiedID *int `json:"lastModifiedId" column:"last_modified_id" sql:"BIGINT"`
	ParentID       *int `json:"parentId"       column:"parent_id" sql:"BIGINT"`

	// ─────── Purchase Info ───────
	PurchaseDate           *time.Time `json:"purchaseDate"            column:"purchase_date" sql:"TIMESTAMPTZ"`
	PurchaseCost           *float64   `json:"purchaseCost"            column:"purchase_cost" sql:"FLOAT"`
	PurchaseCurrency       *string    `json:"purchaseCurrency"        column:"purchase_currency" sql:"VARCHAR"`
	PurchaseWarrantyYears  *int       `json:"purchaseWarrantyDuration" column:"purchase_warranty_duration" sql:"INTEGER"`
	PurchaseCurrencySymbol *string    `json:"purchaseCurrencySymbol"  column:"purchase_currency_symbol" sql:"VARCHAR"`

	// ─────── Metadata ───────
	IsDeleted     bool       `json:"isDeleted"      column:"is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	LastScannedAt *time.Time `json:"lastScannedAt"  column:"last_scanned_at" sql:"TIMESTAMPTZ DEFAULT NOW()"`
	CreatedAt     time.Time  `json:"createdAt"      column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time  `json:"updatedAt"      column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
