package model

import (
	"time"

	util_enum "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

type QR struct {
	ID int `json:"id" column:"qrs.id"`
	// unique
	Code           string                 `json:"code" column:"qrs.code"`
	ShortCode      string                 `json:"shortCode" column:"qrs.short_code"`
	Status         util_enum.TagStatus    `json:"status" column:"qrs.status"`
	LastScanModule util_enum.ModuleAccess `json:"lastScanModule" column:"qrs.last_scan_module"`

	// Nullable fields
	ReportedByID *int       `json:"reportedById" column:"qrs.reported_by_id"`
	ReportedAt   *time.Time `json:"reportedAt" column:"qrs.reported_at"`

	CreatedAt time.Time `json:"createdAt" column:"qrs.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"qrs.updated_at"`
}
func (r *QR) GetColumns() []string {
	return []string{
		"id",
		"code",
		"short_code",
		"status",
		"last_scan_module",
		"reported_by_id",
		"reported_at",
		"created_at",
		"updated_at",
	}
}
