package model

import (
	"time"
)

// Asset
// - DepreciationGroupID:
// For future reference, this is used for the basic version.
// Basic version bypass the creation of depreciation group through asset name, so we set it immediately to the asset.
type Asset struct {
	ID                     int               `json:"_id"                      column:"assets.id"`
	AssetIDFull            string            `json:"assetIdFull"              column:"assets.asset_id_full"`
	AssetID                string            `json:"assetId"                  column:"assets.asset_id"`
	NameID                 int               `json:"nameId"                   column:"assets.name_id"`
	NameWithSequence       string            `json:"nameWithSequence"         column:"assets.name_with_sequence"`
	SequenceNumber         int               `json:"sequenceNumber"           column:"assets.sequence_number"`
	AliasName              *string           `json:"aliasName"                column:"assets.alias_name"`
	AliasCode              *string           `json:"aliasCode"                column:"assets.alias_code"`
	QrID                   *int              `json:"qrId"                     column:"assets.qr_id"`
	RfidID                 *int              `json:"rfidId"                   column:"assets.rfid_id"`
	Status                 AssetStatus       `json:"status"                   column:"assets.status"`
	FirstImage             *string           `json:"firstImage"               column:"assets.first_image"`
	SecondImage            *string           `json:"secondImage"              column:"assets.second_image"`
	SetDefault             AssetDefaultImage `json:"setDefault"               column:"assets.set_default"`
	AssetImage             *string           `json:"assetImage"               column:"assets.asset_image"                special:"generated"`
	PreviousStatus         *string           `json:"previousStatus"           column:"assets.previous_status"`
	TagType                TagType           `json:"tagType"                  column:"assets.tag_type"                   special:"generated"`
	GroupID                int               `json:"groupId"                  column:"assets.group_id"`
	ModelID                int               `json:"modelId"                  column:"assets.model_id"`
	AssignedToID           *int              `json:"assignedToId"             column:"assets.assigned_to_id"`
	BorrowerID             *int              `json:"borrowerId"               column:"assets.borrower_id"`
	LastModifiedID         *int              `json:"lastModifiedId"           column:"assets.last_modified_id"`
	ParentID               *int              `json:"parentId"                 column:"assets.parent_id"`
	DepreciationGroupID    *int              `json:"depreciationGroupId"      column:"assets.depreciation_group_id"`
	PurchaseDate           *time.Time        `json:"purchaseDate"             column:"assets.purchase_date"`
	PurchaseCost           *float64          `json:"purchaseCost"             column:"assets.purchase_cost"`
	PurchaseCurrency       *string           `json:"purchaseCurrency"         column:"assets.purchase_currency"`
	PurchaseWarrantyYears  *int              `json:"purchaseWarrantyDuration" column:"assets.purchase_warranty_duration"`
	PurchaseCurrencySymbol *string           `json:"purchaseCurrencySymbol"   column:"assets.purchase_currency_symbol"`
	IsDeleted              bool              `json:"isDeleted"                column:"assets.is_deleted"`
	LastScannedAt          *time.Time        `json:"lastScannedAt"            column:"assets.last_scanned_at"`
	CreatedAt              time.Time         `json:"createdAt"                column:"assets.created_at"`
	UpdatedAt              time.Time         `json:"updatedAt"                column:"assets.updated_at"`
}
