package new_model

import (
	"time"
)

type Tracking struct {
	ID              int                `json:"_id"             column:"trackings.id"`
	Asset           int                `json:"asset"           column:"trackings.asset"`
	ReportedByID    int                `json:"reportedById"    column:"trackings.reported_by_id"`
	DamageImage     string             `json:"damageImage"     column:"trackings.damage_image"`
	PreviousStatus  AssetStatus        `json:"previousStatus"  column:"trackings.previous_status"`
	TransactionID   string             `json:"transactionId"   column:"trackings.transaction_id"`
	Status          TrackingStatus     `json:"status"          column:"trackings.status"`
	Condition       *TrackingCondition `json:"condition"       column:"trackings.condition"`
	Notes           string             `json:"notes"           column:"trackings.notes"`
	Address         string             `json:"address"         column:"trackings.address"`
	Coordinate      string             `json:"coordinate"      column:"trackings.coordinate"`
	DownTimeSeconds float64            `json:"downTimeSeconds" column:"trackings.down_time_seconds"`
	ReportSource    AssetReportSource  `json:"reportSource"    column:"trackings.report_source"`
	ApprovalRound   int                `json:"approvalRound"   column:"trackings.approval_round"`
	FoundAt         *time.Time         `json:"foundAt"         column:"trackings.found_at"`
	FoundBy         *int               `json:"foundBy"         column:"trackings.found_by_id"`
	FoundGroup      *int               `json:"found_group"     column:"trackings.found_group_id"`
	CreatedAt       time.Time          `json:"createdAt"       column:"trackings.created_at"`
	UpdatedAt       time.Time          `json:"updatedAt"       column:"trackings.updated_at"`
}
