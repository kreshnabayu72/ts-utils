package new_model

import (
	"time"
)

type Rfid struct {
	ID              int         `json:"_id"             column:"rfids.id"`
	Code            string      `json:"code"            column:"rfids.code"`
	ShortCode       string      `json:"shortCode"       column:"rfids.short_code"`
	Status          TagStatus   `json:"status"          column:"rfids.status"`
	LastAssetStatus AssetStatus `json:"lastAssetStatus" column:"rfids.asset_status"`
	HolderId        *int        `json:"holderId"        column:"rfids.holder_id"`
	PreviousAsset   *int        `json:"previousAsset"        column:"rfids.previous_asset"`
	ToBeReturnedAt  *time.Time  `json:"toBeReturnedAt"  column:"rfids.to_be_returned_at"`
	ReportedById    *int        `json:"reportedById"    column:"rfids.reported_by_id"`
	ReportedAt      *time.Time  `json:"reportedAt"      column:"rfids.reported_at"`
	CreatedAt       time.Time   `json:"createdAt"       column:"rfids.created_at"`
	UpdatedAt       time.Time   `json:"updatedAt"       column:"rfids.updated_at"`
}
