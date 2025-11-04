package new_model

import (
	"time"
)

type SampStrategicRiskRegisterUser struct {
	ID int `json:"id" column:"samp_strategic_risk_register_users.id"`

	SampStrategicRiskRegisterId int `json:"sampStrategicRiskRegisterId" column:"samp_strategic_risk_register_users.samp_strategic_risk_register_id"` // Foreign Key to samp_current_versions.id
	UserId                      int `json:"userId" column:"samp_strategic_risk_register_users.user_id"`

	CreatedAt time.Time `json:"createdAt" column:"samp_strategic_risk_register_users.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_strategic_risk_register_users.updated_at"`
}
