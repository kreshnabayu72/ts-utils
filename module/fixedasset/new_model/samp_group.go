package new_model

import (
	"time"
)

type SampGroup struct {
	ID      int `json:"id" column:"samp_groups.id"`
	SampID  int `json:"sampId" column:"samp_groups.samp_id"` // Foreign Key to samp_current_versions.id
	GroupID int `json:"groupId" column:"samp_groups.groups_id"`

	CreatedAt time.Time `json:"createdAt" column:"samp_groups.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_groups.updated_at"`
}
