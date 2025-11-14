package new_model

import (
	"time"
)

type RiskRegisterClosed struct {
	ID int `json:"id" column:"risk_registers_closed.id"`

	RiskID             string    `json:"riskId" column:"risk_registers_closed.risk_id"`
	Risk               string    `json:"risk" column:"risk_registers_closed.risk"`
	Cause              string    `json:"cause" column:"risk_registers_closed.cause"`
	Impact             string    `json:"impact" column:"risk_registers_closed.impact"`
	Category           string    `json:"category" column:"risk_registers_closed.category"`
	LikelihoodSeverity int       `json:"likelihoodSeverity" column:"risk_registers_closed.likelihood_severity"`
	ImpactSeverity     int       `json:"impactSeverity" column:"risk_registers_closed.impact_severity"`
	RiskLevel          string    `json:"riskLevel" column:"risk_registers_closed.risk_level"`
	Mitigation         string    `json:"mitigation" column:"risk_registers_closed.mitigation"`
	TargetDate         time.Time `json:"targetDate" column:"risk_registers_closed.target_date"`
	Status             string    `json:"status" column:"risk_registers_closed.status"`
	PicName            string    `json:"picName" column:"risk_registers_closed.pic_name"`

	CreatedAt time.Time `json:"createdAt" column:"risk_registers_closed.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"risk_registers_closed.updated_at"`
}
