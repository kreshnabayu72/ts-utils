package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Brand struct {
	ID        int        `json:"_id"       column:"brands.id"         sql:"BIGINT PRIMARY KEY"`
	Name      string     `json:"name"      column:"brands.name"       sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(name) > 0)"`
	IsDeleted bool       `json:"isDeleted" column:"brands.is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt *time.Time `json:"deletedAt" column:"brands.deleted_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt time.Time  `json:"createdAt" column:"brands.created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time  `json:"updatedAt" column:"brands.updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Brand) TableName() string {
	return "brands"
}

func (m *Brand) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("idx_brands_name", "name", true, "is_deleted = FALSE")
}
