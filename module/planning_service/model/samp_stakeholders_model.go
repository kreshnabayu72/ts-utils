package new_model

import (
	"time"
)

type SampStakeholders struct {
	ID int `json:"id" column:"samp_stakeholders.id"`

	SampID int `json:"sampId" column:"samp_stakeholders.samp_id"`

	Name         string `json:"name" column:"samp_stakeholders.name"`
	Role         string `json:"role" column:"samp_stakeholders.role"`
	Expectations string `json:"expectations" column:"samp_stakeholders.expectations"`

	CreatedAt time.Time `json:"createdAt" column:"samp_stakeholders.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_stakeholders.updated_at"`
}
