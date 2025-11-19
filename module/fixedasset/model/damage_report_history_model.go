package model

import (
	"time"
)

type DamageReportHistory struct {
	ID               int                `json:"_id"                column:"damage_report_histories.id"`
	Asset            int                `json:"asset"              column:"damage_report_histories.asset"`
	AssetID          string             `json:"assetId"            column:"damage_report_histories.asset_id"`
	AliasCode        *string            `json:"aliasCode"                column:"damage_report_histories.alias_code"`
	TagType          *TagType           `json:"tagType"            column:"damage_report_histories.tag_type"`
	Name             string             `json:"name"               column:"damage_report_histories.name"`
	Brand            string             `json:"brand"              column:"damage_report_histories.brand"`
	Model            string             `json:"model"              column:"damage_report_histories.model"`
	NameWithSequence string             `json:"nameWithSequence"   column:"damage_report_histories.name_with_sequence"`
	SequenceNumber   int                `json:"sequenceNumber"     column:"damage_report_histories.sequence_number"`
	Group            string             `json:"group"              column:"damage_report_histories.group"`
	GroupID          int                `json:"groupId"            column:"damage_report_histories.group_id"`
	ReportedBy       string             `json:"reportedBy"         column:"damage_report_histories.reported_by"`
	DamageImage      string             `json:"damageImage"        column:"damage_report_histories.damage_image"`
	Status           DamageReportStatus `json:"status"             column:"damage_report_histories.status"`
	Notes            string             `json:"notes"              column:"damage_report_histories.notes"`
	RepairedAt       time.Time          `json:"repairedAt"         column:"damage_report_histories.repaired_at"`
	CreatedAt        time.Time          `json:"createdAt"          column:"damage_report_histories.created_at"`
	UpdatedAt        time.Time          `json:"updatedAt"          column:"damage_report_histories.updated_at"`
}
