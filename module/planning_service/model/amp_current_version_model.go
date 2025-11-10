package new_model

import (
	"time"
)

type AmpCurrentVersion struct {
	ID int `json:"id" column:"amp_current_versions.id"`

	IsSubmitted bool `json:"isSubmitted" column:"amp_current_versions.is_submitted"`
	SampID      int  `json:"sampId" column:"amp_current_versions.samp_id"`
	GroupID     int  `json:"groupId" column:"amp_current_versions.group_id"`

	// Need pointers, after init (before submit), the values are null
	AmpID   *string  `json:"ampId" column:"amp_current_versions.amp_id"`
	Name    *string  `json:"name" column:"amp_current_versions.name"`
	Version *float64 `json:"version" column:"amp_current_versions.version"`
	Status  *string  `json:"status" column:"amp_current_versions.status"`

	ModifiedByID int `json:"modifiedById" column:"amp_current_versions.modified_by_id"`

	CreatedAt time.Time `json:"createdAt" column:"amp_current_versions.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"amp_current_versions.updated_at"`
}
