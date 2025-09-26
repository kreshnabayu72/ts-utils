package new_model

import (
	"time"
)

type Brand struct {
	ID        int        `json:"_id"       column:"brands.id"`
	Name      string     `json:"name"      column:"brands.name"`
	IsDeleted bool       `json:"isDeleted" column:"brands.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"brands.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"brands.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"brands.updated_at"`
}
