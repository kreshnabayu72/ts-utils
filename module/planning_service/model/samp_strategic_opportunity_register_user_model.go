package new_model

import (
	"time"
)

type SampStrategicOpportunityRegisterUser struct {
	ID int `json:"id" column:"samp_strategic_opportunity_register_users.id"`

	SampStrategicOpportunityRegisterId int `json:"sampStrategicOpportunityRegisterId" column:"samp_strategic_opportunity_register_users.samp_strategic_opportunity_register_id"` // Foreign Key to samp_current_versions.id
	UserId                             int `json:"userId" column:"samp_strategic_opportunity_register_users.user_id"`

	CreatedAt time.Time `json:"createdAt" column:"samp_strategic_opportunity_register_users.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_strategic_opportunity_register_users.updated_at"`
}
