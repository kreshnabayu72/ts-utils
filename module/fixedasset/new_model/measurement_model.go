package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Measurement struct {
	ID        int        `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Name      string     `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	IsDeleted bool       `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt *time.Time `json:"deletedAt" column:"deleted_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt time.Time  `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time  `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Measurement) TableName() string {
	return "measurements"
}

func (m *Measurement) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("measurements_name_uindex", "name", true, "is_deleted = FALSE")
}
