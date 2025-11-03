package new_model

import (
	"time"
)

type SampInternalContext struct {
	ID     int `json:"id" column:"samp_internal_contexts.id"`
	SampID int `json:"sampId" column:"samp_internal_contexts.samp_id"` // Foreign Key to samp_current_versions.id

	Name    string `json:"name" column:"samp_internal_contexts.name"`
	Context string `json:"context" column:"samp_internal_contexts.context"`

	CreatedAt time.Time `json:"createdAt" column:"samp_internal_contexts.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_internal_contexts.updated_at"`
}
