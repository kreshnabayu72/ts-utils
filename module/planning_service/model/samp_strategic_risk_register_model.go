package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/new_model"
)

type SampStrategicRiskRegister struct {
	ID int `json:"id" column:"samp_strategic_risk_registers.id"`

	SampId             int                 `json:"sampId" column:"samp_strategic_risk_registers.samp_id"` // Foreign Key to samp_current_versions.id
	Risk               string              `json:"risk" column:"samp_strategic_risk_registers.risk"`
	Cause              string              `json:"cause" column:"samp_strategic_risk_registers.cause"`
	Impact             string              `json:"impact" column:"samp_strategic_risk_registers.impact"`
	Mitigation         string              `json:"mitigation" column:"samp_strategic_risk_registers.mitigation"`
	Category           string              `json:"category" column:"samp_strategic_risk_registers.category"`
	LikelihoodSeverity int                 `json:"likelihoodSeverity" column:"samp_strategic_risk_registers.likelihoodSeverity"`
	ImpactSeverity     int                 `json:"impactSeverity" column:"samp_strategic_risk_registers.impactSeverity"`
	Timeline           new_model.TimeRange `json:"timeline" column:"samp_strategic_risk_registers.timeline"`

	CreatedAt time.Time `json:"createdAt" column:"samp_strategic_risk_registers.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_strategic_risk_registers.updated_at"`
}
