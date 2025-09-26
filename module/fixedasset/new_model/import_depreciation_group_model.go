package new_model

import (
	"time"
)

type ImportDepreciationGroup struct {
	ID         int       `json:"_id"        column:"import_depreciation_groups.id"`
	Key        int       `json:"key"        column:"import_depreciation_groups.key"         special:"generated"`
	Name       string    `json:"name"       column:"import_depreciation_groups.name"`
	UsefulLife int       `json:"usefulLife" column:"import_depreciation_groups.useful_life"`
	CreatedAt  time.Time `json:"createdAt"  column:"import_depreciation_groups.created_at"`
	UpdatedAt  time.Time `json:"updatedAt"  column:"import_depreciation_groups.updated_at"`
}
