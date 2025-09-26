package new_model

import (
	"time"
)

type Category struct {
	ID        int        `json:"_id"       column:"categories.id"`
	Name      string     `json:"name"      column:"categories.name"`
	Code      string     `json:"code"      column:"categories.code"`
	ParentId  int        `json:"parentId"  column:"categories.parent_id"`
	IsDeleted bool       `json:"isDeleted" column:"categories.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"categories.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"categories.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"categories.updated_at"`
}
