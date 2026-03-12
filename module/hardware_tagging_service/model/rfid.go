package model

import (
	"time"

	util_enum "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

type RFID struct {
	ID int `json:"_id" column:"rfids.id"`
	// unique
	Status              util_enum.TagStatus    `json:"status" column:"rfids.status"`
	LastScanModule      util_enum.ModuleAccess `json:"lastScanModule" column:"rfids.last_scan_module"`
	ToBeReturnedAssetID *int                   `json:"toBeReturnedAssetId" column:"rfids.to_be_returned_asset_id"`
	ToBeReturnedAt      *time.Time             `json:"toBeReturnedAt" column:"rfids.to_be_returned_at"`
	ReportedById        *int                   `json:"reportedById" column:"rfids.reported_by_id"`
	ReportedAt          *time.Time             `json:"reportedAt" column:"rfids.reported_at"`
	CreatedAt           *time.Time             `json:"createdAt" column:"rfids.created_at"`
	UpdatedAt           *time.Time             `json:"updatedAt" column:"rfids.updated_at"`
}
func (r *RFID) GetColumns() []string {
	return []string{
		"id",
		"status",
		"last_scan_module",
		"to_be_returned_asset_id",
		"to_be_returned_at",
		"reported_by_id",
		"reported_at",
		"created_at",
		"updated_at",
	}
}
