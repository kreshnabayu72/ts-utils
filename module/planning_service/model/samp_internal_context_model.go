package new_model

import (
	"time"
)

type SampInternalContext struct {
	ID            int `json:"id" column:"samp_internal_contexts.id"`
	SampID        int `json:"sampId" column:"samp_internal_contexts.samp_id"`                // Foreign Key to samp
	SampHistoryID int `json:"sampHistoryId" column:"samp_internal_contexts.samp_history_id"` // Foreign Key to samp history

	Name    string `json:"name" column:"samp_internal_contexts.name"`
	Context string `json:"context" column:"samp_internal_contexts.context"`

	IsDeleted bool       `json:"isDeleted"        column:"samp_internal_contexts.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"samp_internal_contexts.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"samp_internal_contexts.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_internal_contexts.updated_at"`
}
