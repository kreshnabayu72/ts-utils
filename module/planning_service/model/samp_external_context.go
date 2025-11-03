package new_model

import (
	"time"
)

type SampExternalContext struct {
	ID     int `json:"id" column:"samp_external_contexts.id"`
	SampID int `json:"sampId" column:"samp_external_contexts.samp_id"` // Foreign Key to samp_current_versions.id

	Name    string `json:"name" column:"samp_external_contexts.name"`
	Context string `json:"context" column:"samp_external_contexts.context"`

	CreatedAt time.Time `json:"createdAt" column:"samp_external_contexts.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_external_contexts.updated_at"`
}
