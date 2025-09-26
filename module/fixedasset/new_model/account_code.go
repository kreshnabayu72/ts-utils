package new_model

import (
	"time"
)

type AccountCode struct {
	ID        int        `json:"_id"       column:"account_codes.id"`
	Name      string     `json:"name"      column:"account_codes.name"`
	Code      string     `json:"code"      column:"account_codes.code"`
	IsDeleted bool       `json:"isDeleted" column:"account_codes.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"account_codes.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"account_codes.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"account_codes.updated_at"`
}
