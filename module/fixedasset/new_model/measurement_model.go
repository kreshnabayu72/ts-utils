package new_model

import (
	"time"
)

type Measurement struct {
	ID        int        `json:"_id"                  column:"id"`
	Name      string     `json:"name"                 column:"name"`
	IsDeleted bool       `json:"isDeleted"            column:"is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"deleted_at"`
	CreatedAt time.Time  `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt"            column:"updated_at"`
}
