package new_model

import (
	"time"
)

type RepairTicketingHistory struct {
	ID                         int                     `json:"_id"                        column:"repair_ticketing_histories.id"`
	Asset                      int                     `json:"asset"                      column:"repair_ticketing_histories.asset"`
	ReportedByFullName         string                  `json:"reportedByFullName"         column:"repair_ticketing_histories.reported_by_full_name"`
	TagType                    TagType                 `json:"tagType"                    column:"repair_ticketing_histories.tag_type"`
	AssetNameWithSequence      string                  `json:"assetNameWithSequence"      column:"repair_ticketing_histories.asset_name_with_sequence"`
	AssetID                    string                  `json:"assetId"             column:"repair_ticketing_histories.asset_id"`
	AssetAliasCode             *string                 `json:"assetAliasCode"             column:"repair_ticketing_histories.asset_alias_code"`
	AssetBrand                 string                  `json:"assetBrand"                 column:"repair_ticketing_histories.asset_brand"`
	AssetModel                 string                  `json:"assetModel"                 column:"repair_ticketing_histories.asset_model"`
	AssetGroupName             string                  `json:"assetGroup"                 column:"repair_ticketing_histories.asset_group_name"`
	AssetGroupId               int                     `json:"assetGroupId"               column:"repair_ticketing_histories.asset_group_id"`
	DamageImage                string                  `json:"damageImage"                column:"repair_ticketing_histories.damage_image"`
	PreviousStatus             AssetStatus             `json:"previousStatus"             column:"repair_ticketing_histories.previous_status"`
	TransactionID              string                  `json:"transactionId"              column:"repair_ticketing_histories.transaction_id"`
	Status                     RepairTicketingStatus   `json:"status"                     column:"repair_ticketing_histories.status"`
	PreviousMethod             []RepairTicketingMethod `json:"previousMethod"             column:"repair_ticketing_histories.previous_method"`
	ReportNotes                string                  `json:"reportNotes"                column:"repair_ticketing_histories.report_notes"`
	RepairNotes                string                  `json:"repairNotes"                column:"repair_ticketing_histories.repair_notes"`
	IsRepairSuccess            *bool                   `json:"isRepairSuccess"            column:"repair_ticketing_histories.is_repair_success"`
	DownTimeSeconds            float64                 `json:"downTimeSeconds"            column:"repair_ticketing_histories.down_time_seconds"`
	ReportSource               AssetReportSource       `json:"reportSource"               column:"repair_ticketing_histories.report_source"`
	Details                    string                  `json:"details"                    column:"repair_ticketing_histories.details"`
	HandledBy                  RepairTicketingMethod   `json:"handledBy"                  column:"repair_ticketing_histories.handled_by"`
	ApprovalRound              int                     `json:"approvalRound"              column:"repair_ticketings.approval_round"`
	AdditionalCostTotalPrice   float32                 `json:"additionalCostTotalPrice"   column:"repair_ticketing_histories.additional_cost_total_price"`
	AdditionalCostCurrency     string                  `json:"additionalCostCurrency"     column:"repair_ticketing_histories.additional_cost_currency"`
	AdditionalCostPaymentProof string                  `json:"additionalCostPaymentProof" column:"repair_ticketing_histories.additional_cost_payment_proof"`
	Address                    string                  `json:"address"                    column:"repair_ticketing_histories.address"`
	Coordinate                 string                  `json:"coordinate"                 column:"repair_ticketing_histories.coordinate"`
	StaffFullName              string                  `json:"staffFullName"              column:"repair_ticketing_histories.staff_full_name"`
	ServiceCenterName          *string                 `json:"serviceCenterName"          column:"repair_ticketing_histories.service_center_name"`
	RepairStartAt              *time.Time              `json:"repairStartAt"              column:"repair_ticketing_histories.repair_start_at"`
	CreatedAt                  time.Time               `json:"createdAt"                  column:"repair_ticketing_histories.created_at"`
	UpdatedAt                  time.Time               `json:"updatedAt"                  column:"repair_ticketing_histories.updated_at"`
}
