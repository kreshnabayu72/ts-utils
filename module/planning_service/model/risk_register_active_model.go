package new_model

import (
	"time"
)

type RiskRegisterActive struct {
	ID int `json:"id" column:"risk_registers_active.id"`

	RiskID             string    `json:"riskId" column:"risk_registers_active.risk_id"`
	Risk               string    `json:"risk" column:"risk_registers_active.risk"`
	Cause              string    `json:"cause" column:"risk_registers_active.cause"`
	Impact             string    `json:"impact" column:"risk_registers_active.impact"`
	Category           string    `json:"category" column:"risk_registers_active.category"`
	LikelihoodSeverity int       `json:"likelihoodSeverity" column:"risk_registers_active.likelihood_severity"`
	ImpactSeverity     int       `json:"impactSeverity" column:"risk_registers_active.impact_severity"`
	RiskLevel          string    `json:"riskLevel" column:"risk_registers_active.risk_level"`
	Mitigation         string    `json:"mitigation" column:"risk_registers_active.mitigation"`
	TargetDate         time.Time `json:"targetDate" column:"risk_registers_active.target_date"`
	Status             string    `json:"status" column:"risk_registers_active.status"`
	PicID              int       `json:"picId" column:"risk_registers_active.pic_id"`

	CreatedAt time.Time `json:"createdAt" column:"risk_registers_active.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"risk_registers_active.updated_at"`

	IsDeleted bool       `json:"isDeleted" column:"risk_registers_active.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"risk_registers_active.deleted_at"`
}
