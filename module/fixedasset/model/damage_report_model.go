package model

import (
	"time"
)

type DamageReport struct {
	ID             int                `json:"_id"            column:"damage_reports.id"`
	Asset          int                `json:"asset"          column:"damage_reports.asset"`
	ReportedById   int                `json:"reportedById"   column:"damage_reports.reported_by_id"`
	DamageImage    string             `json:"damageImage"    column:"damage_reports.damage_image"`
	Status         DamageReportStatus `json:"status"         column:"damage_reports.status"`
	PreviousStatus AssetStatus        `json:"previousStatus" column:"damage_reports.previous_status"`
	ReportSource   AssetReportSource  `json:"reportSource"   column:"damage_reports.report_source"`
	Notes          string             `json:"notes"          column:"damage_reports.notes"`
	RepairedAt     *time.Time         `json:"repairedAt"     column:"damage_reports.repaired_at"`
	CreatedAt      time.Time          `json:"createdAt"      column:"damage_reports.created_at"`
	UpdatedAt      time.Time          `json:"updatedAt"      column:"damage_reports.updated_at"`
}
