package new_model

import (
	"time"
)

type SampAmObjectiveUser struct {
	ID int `json:"id" column:"samp_am_objective_users.id"`

	SampAmObjectiveID int `json:"sampAmObjectiveId" column:"samp_am_objective_users.samp_am_objective_id"` // Foreign Key to samp_current_versions.id
	UserId            int `json:"sampId" column:"samp_am_objective_users.user_id"`                         // Foreign Key to samp_current_versions.id

	IsDeleted bool       `json:"isDeleted"        column:"samp_am_objective_users.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"samp_am_objective_users.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"samp_am_objective_users.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_am_objective_users.updated_at"`
}
