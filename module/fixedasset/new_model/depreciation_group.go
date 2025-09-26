package new_model

import (
	"time"
)

type DepreciationGroup struct {
	ID                  int        `json:"_id"                 column:"depreciation_groups.id"`
	Name                string     `json:"name"                column:"depreciation_groups.name"`
	UsefulLife          int        `json:"usefulLife"          column:"depreciation_groups.useful_life"`
	StraightLineRate    float64    `json:"straightLineRate"    column:"depreciation_groups.straight_line_rate"    special:"generated"`
	DoubleDecliningRate float64    `json:"doubleDecliningRate" column:"depreciation_groups.double_declining_rate" special:"generated"`
	IsDeleted           bool       `json:"isDeleted"           column:"depreciation_groups.is_deleted"`
	DeletedAt           *time.Time `json:"deletedAt"           column:"depreciation_groups.deleted_at"`
	CreatedAt           time.Time  `json:"createdAt"           column:"depreciation_groups.created_at"`
	UpdatedAt           time.Time  `json:"updatedAt"           column:"depreciation_groups.updated_at"`
}
