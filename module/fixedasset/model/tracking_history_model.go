package model

import (
	"time"
)

type TrackingHistory struct {
	ID                    int               `json:"_id"                   column:"tracking_histories.id"`
	Asset                 int               `json:"asset"                 column:"tracking_histories.asset"`
	TagType               TagType           `json:"tagType"               column:"tracking_histories.tag_type"`
	ReportedByFullName    string            `json:"reportedByFullName"    column:"tracking_histories.reported_by_full_name"`
	AssetNameWithSequence string            `json:"assetNameWithSequence" column:"tracking_histories.asset_name_with_sequence"`
	AssetID               string            `json:"assetId"             column:"tracking_histories.asset_id"`
	AssetAliasCode        *string           `json:"assetAliasCode"        column:"tracking_histories.asset_alias_code"`
	AssetBrand            string            `json:"assetBrand"            column:"tracking_histories.asset_brand"`
	AssetModel            string            `json:"assetModel"            column:"tracking_histories.asset_model"`
	AssetGroupName        string            `json:"assetGroup"            column:"tracking_histories.asset_group_name"`
	AssetGroupId          int               `json:"assetGroupId"          column:"tracking_histories.asset_group_id"`
	DamageImage           string            `json:"damageImage"           column:"tracking_histories.damage_image"`
	PreviousStatus        AssetStatus       `json:"previousStatus"        column:"tracking_histories.previous_status"`
	TransactionID         string            `json:"transactionId"         column:"tracking_histories.transaction_id"`
	Status                TrackingStatus    `json:"status"                column:"tracking_histories.status"`
	Condition             TrackingCondition `json:"condition"             column:"tracking_histories.condition"`
	Notes                 string            `json:"notes"                 column:"tracking_histories.notes"`
	Address               string            `json:"address"               column:"tracking_histories.address"`
	Coordinate            string            `json:"coordinate"            column:"tracking_histories.coordinate"`
	DownTimeSeconds       float64           `json:"downTimeSeconds"       column:"tracking_histories.down_time_seconds"`
	ReportSource          AssetReportSource `json:"reportSource"          column:"tracking_histories.report_source"`
	ApprovalRound         int               `json:"approvalRound"         column:"tracking_histories.approval_round"`
	FoundAt               time.Time         `json:"foundAt"               column:"tracking_histories.found_at"`
	FoundById             int               `json:"foundBy"               column:"tracking_histories.found_by_id"`
	FoundByFullName       string            `json:"foundByFullName"       column:"tracking_histories.found_by_full_name"`
	FoundGroupId          int               `json:"foundGroupId"          column:"tracking_histories.found_group_id"`
	FoundGroupName        string            `json:"foundGroupName"        column:"tracking_histories.found_group_name"`
	CreatedAt             time.Time         `json:"createdAt"             column:"tracking_histories.created_at"`
	UpdatedAt             time.Time         `json:"updatedAt"             column:"tracking_histories.updated_at"`
}
