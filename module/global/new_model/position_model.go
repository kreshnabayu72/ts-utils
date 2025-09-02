package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	global_migration "gitea.qwertysystem.net/BETS/ts-utils/module/global/new_model/migration"
)

type Position struct {
	ID        int       `json:"_id" column:"positions.id" sql:"BIGINT PRIMARY KEY"`
	Name      string    `json:"name" column:"positions.name" sql:"VARCHAR NOT NULL"`
	CreatedAt time.Time `json:"createdAt" column:"positions.created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"positions.updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Position) TableName() string {
	return "positions"
}

func (m *Position) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(global_migration.PositionPostQuery)
}
