package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Division struct {
	ID        int       `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Name      string    `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (m *Division) TableName() string {
	return "divisions"
}

func (m *Division) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}
