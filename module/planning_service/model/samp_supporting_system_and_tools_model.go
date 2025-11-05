package new_model

import (
	"time"
)

type SampSupportingSystemAndTools struct {
	ID int `json:"id" column:"samp_supporting_system_and_tools.id"`

	SampID   int    `json:"sampId" column:"samp_supporting_system_and_tools.samp_id"`
	Name     string `json:"name" column:"samp_supporting_system_and_tools.name"`
	Type     string `json:"type" column:"samp_supporting_system_and_tools.type"`
	Function string `json:"function" column:"samp_supporting_system_and_tools.function"`
	Notes    string `json:"notes" column:"samp_supporting_system_and_tools.notes"`

	IsDeleted bool       `json:"isDeleted"        column:"samp_supporting_system_and_tools.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"samp_supporting_system_and_tools.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"samp_supporting_system_and_tools.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_supporting_system_and_tools.updated_at"`
}
