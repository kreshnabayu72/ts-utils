package model

import (
	"time"

	util_enum "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

type NFC struct {
	ID int `json:"id" column:"nfcs.id"`
	// unique
	Status         util_enum.TagStatus    `json:"status" column:"nfcs.status"`
	LastScanModule util_enum.ModuleAccess `json:"lastScanModule" column:"nfcs.last_scan_module"`

	ToBeReturnedAssetID  *int       `json:"toBeReturnedAssetId" column:"nfcs.to_be_returned_asset_id"`
	ToBeReturnedHolderID *int       `json:"toBeReturnedHolderId" column:"nfcs.to_be_returned_holder_id"`
	ToBeReturnedAt       *time.Time `json:"toBeReturnedAt" column:"nfcs.to_be_returned_at"`
	ReportedByID         *int       `json:"reportedById" column:"nfcs.reported_by_id"`
	ReportedAt           *time.Time `json:"reportedAt" column:"nfcs.reported_at"`

	CreatedAt time.Time `json:"createdAt" column:"nfcs.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"nfcs.updated_at"`
}
func (r *NFC) GetColumns() []string {
	return []string{
		"id",
		"status",
		"last_scan_module",
		"to_be_returned_asset_id",
		"to_be_returned_holder_id",
		"to_be_returned_at",
		"reported_by_id",
		"reported_at",
		"created_at",
		"updated_at",
	}
}
