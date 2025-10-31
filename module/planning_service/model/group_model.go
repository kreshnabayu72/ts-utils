package new_model

import (
	"time"
)

type Group struct {
	ID       int    `json:"id"                                 column:"groups.id"`
	Name     string `json:"name"                               column:"groups.name"`
	FullPath string `json:"fullPath"                               column:"groups.full_path"`

	IsDeleted bool      `json:"isDeleted"                          column:"groups.is_deleted"`
	CreatedAt time.Time `json:"createdAt"                          column:"groups.created_at"`
	UpdatedAt time.Time `json:"updatedAt"                          column:"groups.updated_at"`
}
