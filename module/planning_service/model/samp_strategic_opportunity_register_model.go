package new_model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/common/model"
)

type SampStrategicOpportunityRegister struct {
	ID int `json:"id" column:"samp_strategic_opportunity_registers.id"`

	SampID        int             `json:"sampId" column:"samp_stakeholders.samp_id"`                // Foreign Key to samp
	SampHistoryID int             `json:"sampHistoryId" column:"samp_stakeholders.samp_history_id"` // Foreign Key to samp history
	Opportunity   string          `json:"opportunity" column:"samp_strategic_opportunity_registers.opportunity"`
	Benefit       string          `json:"benefit" column:"samp_strategic_opportunity_registers.benefit"`
	Action        string          `json:"action" column:"samp_strategic_opportunity_registers.action"`
	Timeline      model.TimeRange `json:"timeline" column:"samp_strategic_opportunity_registers.timeline"`

	IsDeleted bool       `json:"isDeleted"        column:"samp_strategic_opportunity_registers.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"samp_strategic_opportunity_registers.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"samp_strategic_opportunity_registers.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_strategic_opportunity_registers.updated_at"`
}
