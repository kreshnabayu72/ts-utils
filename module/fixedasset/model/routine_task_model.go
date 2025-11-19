package model

import (
	"time"
)

type RoutineTask struct {
	ID                         int               `json:"_id"                        column:"routine_tasks.id"`
	ScheduleID                 int               `json:"scheduleId"                 column:"routine_tasks.schedule_id"`
	Name                       string            `json:"name"                       column:"routine_tasks.name"`
	AssetID                    int               `json:"assetId"                    column:"routine_tasks.asset_id"`
	Status                     RoutineTaskStatus `json:"status"                     column:"routine_tasks.status"`
	Duration                   int               `json:"duration"                   column:"routine_tasks.duration"`
	ModifiedByID               int               `json:"modifiedById"               column:"routine_tasks.modified_by_id"`
	ApprovalRound              int               `json:"approvalRound"              column:"routine_tasks.approval_round"`
	Address                    string            `json:"address"                    column:"routine_tasks.address"`
	Coordinate                 string            `json:"coordinate"                 column:"routine_tasks.coordinate"`
	SerialNumber               *string           `json:"serialNumber"               column:"routine_tasks.serial_number"`
	ServiceCenterId            *int              `json:"serviceCenterId"            column:"routine_tasks.service_center_id"`
	StartAt                    time.Time         `json:"startAt"                    column:"routine_tasks.start_at"`
	ReviewedAt                 *time.Time        `json:"reviewedAt"                 column:"routine_tasks.reviewed_at"`
	ReviewedById               *int              `json:"reviewedById"               column:"routine_tasks.reviewed_by_id"`
	Notes                      *string           `json:"notes"                      column:"routine_tasks.notes"`
	AdditionalCostValue        *float64          `json:"additionalCostValue"        column:"routine_tasks.additional_cost_value"`
	AdditionalCostCurrency     *string           `json:"additionalCostCurrency"     column:"routine_tasks.additional_cost_currency"`
	AdditionalCostPaymentProof *string           `json:"additionalCostPaymentProof" column:"routine_tasks.additional_cost_payment_proof"`
	MaintenanceImage           *string           `json:"maintenanceImage"           column:"routine_tasks.maintenance_image"`
	CreatedAt                  time.Time         `json:"createdAt"                  column:"routine_tasks.created_at"`
	UpdatedAt                  time.Time         `json:"updatedAt"                  column:"routine_tasks.updated_at"`
}
