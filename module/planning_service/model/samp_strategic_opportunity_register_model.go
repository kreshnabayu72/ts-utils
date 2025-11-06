package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/new_model"
)

type SampStrategicOpportunityRegister struct {
	ID int `json:"id" column:"samp_strategic_opportunity_registers.id"`

	SampId      int                 `json:"sampId" column:"samp_strategic_opportunity_registers.samp_id"` // Foreign Key to samp_current_versions.id
	Opportunity string              `json:"opportunity" column:"samp_strategic_opportunity_registers.opportunity"`
	Benefit     string              `json:"benefit" column:"samp_strategic_opportunity_registers.benefit"`
	Action      string              `json:"action" column:"samp_strategic_opportunity_registers.action"`
	Timeline    new_model.TimeRange `json:"timeline" column:"samp_strategic_opportunity_registers.timeline"`

	IsDeleted bool       `json:"isDeleted"        column:"samp_strategic_opportunity_registers.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"samp_strategic_opportunity_registers.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"samp_strategic_opportunity_registers.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_strategic_opportunity_registers.updated_at"`
}
