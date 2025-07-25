package new_model

import (
	"time"
)

type DepreciationGroup struct {
	ID                   int        `json:"_id"                  column:"id"`
	Name                 string     `json:"name"                 column:"name"`
	UsefulLife           int        `json:"usefulLife"                 column:"useful_life"`
	StraightLineRate     float64    `json:"straightLineRate"                 column:"straight_line_rate"`
	DoubleDecliningRate  float64    `json:"doubleDecliningRate"                 column:"double_declining_rate"`
	IsDeleted            bool       `json:"isDeleted"            column:"is_deleted"`
	DeleteFromDatabaseAt *time.Time `json:"deleteFromDatabaseAt" column:"delete_from_database_at"`
	CreatedAt            time.Time  `json:"createdAt"            column:"created_at"`
	UpdatedAt            time.Time  `json:"updatedAt"            column:"updated_at"`
}
