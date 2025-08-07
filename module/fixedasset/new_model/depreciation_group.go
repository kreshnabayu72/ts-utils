package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type DepreciationGroup struct {
	ID                  int        `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Name                string     `json:"name" column:"name" sql:"TEXT NOT NULL UNIQUE"`
	UsefulLife          int        `json:"usefulLife" column:"useful_life" sql:"BIGINT NOT NULL"`
	StraightLineRate    float64    `json:"straightLineRate" column:"straight_line_rate" sql:"NUMERIC GENERATED ALWAYS AS (100.0 / useful_life) STORED"`
	DoubleDecliningRate float64    `json:"doubleDecliningRate" column:"double_declining_rate" sql:"NUMERIC GENERATED ALWAYS AS (200.0 / useful_life) STORED"`
	IsDeleted           bool       `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt           *time.Time `json:"deletedAt" column:"deleted_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt           time.Time  `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt           time.Time  `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *DepreciationGroup) TableName() string {
	return "depreciation_groups"
}

func (m *DepreciationGroup) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("depreciation_groups_name_uindex", "name", true, "is_deleted = FALSE")
}
