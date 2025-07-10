package new_model

import "time"

type Asset struct {
	ID int `json:"id" column:"id"`

	AssetID        string `json:"assetId"        column:"asset_id"`
	NameID         int    `json:"nameId"         column:"name_id"`
	SequenceNumber int    `json:"sequenceNumber" column:"sequence_number"`
	AliasName      string `json:"aliasName"      column:"alias_name"`
	QrID           int    `json:"qrId"           column:"qr_id"`
	RfidID         int    `json:"rfidId"         column:"rfid_id"`
	Status         string `json:"status"         column:"status"`
	FirstImage     string `json:"firstImage"     column:"first_image"`
	SecondImage    string `json:"secondImage"    column:"second_image"`
	SetDefault     string `json:"setDefault"     column:"set_default"` // enum: "first_image" | "second_image"
	PreviousStatus string `json:"previousStatus" column:"previous_status"`

	// ─────── References ───────
	GroupID        int `json:"groupId"      column:"group_id"`
	ModelID        int `json:"modelId"      column:"model_id"`
	AssignedToID   int `json:"assignedToId" column:"assigned_to_id"`
	BorrowerID     int `json:"borrowerId"   column:"borrower_id"`
	LastModifiedID int `json:"lastModifiedId" column:"last_modified_id"`
	ParentID       int `json:"parentId"       column:"parent_id"`

	// ─────── Purchase Info ───────
	PurchaseDate          *time.Time `json:"purchaseDate"          column:"purchase_date"`
	PurchaseCost          float64    `json:"purchaseCost"          column:"purchase_cost"`
	PurchaseCurrency      string     `json:"purchaseCurrency"      column:"purchase_currency"`
	PurchaseWarrantyYears int        `json:"purchaseWarrantyYears" column:"purchase_warranty_duration"`

	// ─────── Metadata ───────
	IsDeleted     bool       `json:"isDeleted"      column:"is_deleted"`
	LastScannedAt *time.Time `json:"lastScannedAt"  column:"last_scanned_at"`
	CreatedAt     time.Time  `json:"createdAt"      column:"created_at"`
	UpdatedAt     time.Time  `json:"updatedAt"      column:"updated_at"`
}
