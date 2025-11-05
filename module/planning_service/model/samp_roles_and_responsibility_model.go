package new_model

import (
	"time"
)

type SampRolesAndResponsibilities struct {
	ID int `json:"id" column:"samp_roles_and_responsibilities.id"`

	SampID         int    `json:"sampId" column:"samp_roles_and_responsibilities.samp_id"`
	UserID         int    `json:"userId" column:"samp_roles_and_responsibilities.user_id"`
	Responsibility string `json:"responsibility" column:"samp_roles_and_responsibilities.responsibility"`

	IsDeleted bool       `json:"isDeleted"        column:"samp_roles_and_responsibilities.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"samp_roles_and_responsibilities.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"samp_roles_and_responsibilities.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_roles_and_responsibilities.updated_at"`
}
