package model

import (
	"time"
)

type MissingReportHistory struct {
	ID               int                 `json:"_id"              column:"missing_report_histories.id"`
	Asset            int                 `json:"asset"            column:"missing_report_histories.asset"`
	AssetID          string              `json:"assetId"          column:"missing_report_histories.asset_id"`
	AliasCode        *string             `json:"aliasCode"        column:"missing_report_histories.alias_code"`
	TagType          *TagType            `json:"tagType"          column:"missing_report_histories.tag_type"`
	Name             string              `json:"name"             column:"missing_report_histories.name"`
	Brand            string              `json:"brand"            column:"missing_report_histories.brand"`
	Model            string              `json:"model"            column:"missing_report_histories.model"`
	NameWithSequence string              `json:"nameWithSequence" column:"missing_report_histories.name_with_sequence"`
	SequenceNumber   int                 `json:"sequenceNumber"   column:"missing_report_histories.sequence_number"`
	Image            *string             `json:"image"            column:"missing_report_histories.image"`
	Group            string              `json:"group"            column:"missing_report_histories.group"`
	GroupID          int                 `json:"groupId"          column:"missing_report_histories.group_id"`
	ReportedBy       string              `json:"reportedBy"       column:"missing_report_histories.reported_by"`
	ReportSource     AssetReportSource   `json:"reportSource"     column:"missing_report_histories.report_source"`
	Status           MissingReportStatus `json:"status"           column:"missing_report_histories.status"`
	Notes            string              `json:"notes"            column:"missing_report_histories.notes"`
	FoundAt          time.Time           `json:"foundAt"          column:"missing_report_histories.found_at"`
	CreatedAt        time.Time           `json:"createdAt"        column:"missing_report_histories.created_at"`
	UpdatedAt        time.Time           `json:"updatedAt"        column:"missing_report_histories.updated_at"`
}
