package new_model

import (
	"time"
)

type RepairTicketing struct {
	ID                                 int                     `json:"_id"                        column:"repair_ticketings.id"`
	Asset                              int                     `json:"asset"                      column:"repair_ticketings.asset"`
	ReportedByID                       int                     `json:"reportedById"               column:"repair_ticketings.reported_by_id"`
	RepairImage                        string                  `json:"repairImage"                column:"repair_ticketings.repair_image"`
	DamageImage                        string                  `json:"damageImage"                column:"repair_ticketings.damage_image"`
	PreviousStatus                     AssetStatus             `json:"previousStatus"             column:"repair_ticketings.previous_status"`
	TransactionID                      string                  `json:"transactionId"              column:"repair_ticketings.transaction_id"`
	Status                             RepairTicketingStatus   `json:"status"                     column:"repair_ticketings.status"`
	PreviousMethod                     []RepairTicketingMethod `json:"previousMethod"             column:"repair_ticketings.previous_method"`
	ReportNotes                        string                  `json:"reportNotes"                column:"repair_ticketings.report_notes"`
	RepairNotes                        string                  `json:"repairNotes"                column:"repair_ticketings.repair_notes"`
	IsRepairSuccess                    *bool                   `json:"isRepairSuccess"            column:"repair_ticketings.is_repair_success"`
	DownTimeSeconds                    float64                 `json:"downTimeSeconds"            column:"repair_ticketings.down_time_seconds"`
	ReportSource                       AssetReportSource       `json:"reportSource"               column:"repair_ticketings.report_source"`
	ApprovalRound                      int                     `json:"approvalRound"              column:"repair_ticketings.approval_round"`
	Details                            string                  `json:"details"                    column:"repair_ticketings.details"`
	HandledBy                          RepairTicketingMethod   `json:"handledBy"                  column:"repair_ticketings.handled_by"`
	AdditionalCostTotalPrice           float32                 `json:"additionalCostTotalPrice"   column:"repair_ticketings.additional_cost_total_price"`
	AdditionalCostCurrency             string                  `json:"additionalCostCurrency"     column:"repair_ticketings.additional_cost_currency"`
	AdditionalCostPaymentProof         string                  `json:"additionalCostPaymentProof" column:"repair_ticketings.additional_cost_payment_proof"`
	AdditionalCostPaymentProofFileName string                  `json:"additionalCostPaymentProofFileName" column:"repair_ticketings.additional_cost_payment_proof_file_name"`
	StaffID                            *int                    `json:"staffId"                    column:"repair_ticketings.staff_id"`
	Address                            string                  `json:"address"                    column:"repair_ticketings.address"`
	Coordinate                         string                  `json:"coordinate"                 column:"repair_ticketings.coordinate"`
	ServiceCenterId                    *int                    `json:"serviceCenterId"            column:"repair_ticketings.service_center_id"`
	RepairStartAt                      *time.Time              `json:"repairStartAt"              column:"repair_ticketings.repair_start_at"`
	RepairFinishedAt                   *time.Time              `json:"repairFinishedAt"           column:"repair_ticketings.repair_finished_at"`
	CreatedAt                          time.Time               `json:"createdAt"                  column:"repair_ticketings.created_at"`
	UpdatedAt                          time.Time               `json:"updatedAt"                  column:"repair_ticketings.updated_at"`
}
