package new_model

import (
	"time"
)

type SampStakeholders struct {
	ID int `json:"id" column:"samp_stakeholders.id"`

	SampID        int `json:"sampId" column:"samp_stakeholders.samp_id"`                // Foreign Key to samp
	SampHistoryID int `json:"sampHistoryId" column:"samp_stakeholders.samp_history_id"` // Foreign Key to samp history

	Name         string `json:"name" column:"samp_stakeholders.name"`
	Role         string `json:"role" column:"samp_stakeholders.role"`
	Expectations string `json:"expectations" column:"samp_stakeholders.expectations"`

	IsDeleted bool       `json:"isDeleted"        column:"samp_stakeholders.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"samp_stakeholders.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"samp_stakeholders.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_stakeholders.updated_at"`
}
