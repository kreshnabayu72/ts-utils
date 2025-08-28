package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	global_migration "gitea.qwertysystem.net/BETS/ts-utils/module/global/migration"
)

type Position struct {
	ID        int       `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Name      string    `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Position) TableName() string {
	return "positions"
}

func (m *Position) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.SetPostQuery(global_migration.PositionPostQuery)
}
