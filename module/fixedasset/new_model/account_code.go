package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AccountCode struct {
	ID        int        `json:"_id"       column:"account_codes.id"         sql:"BIGINT PRIMARY KEY"`
	Name      string     `json:"name"      column:"account_codes.name"       sql:"VARCHAR NOT NULL UNIQUE"`
	Code      string     `json:"code"      column:"account_codes.code"       sql:"VARCHAR NOT NULL UNIQUE"`
	IsDeleted bool       `json:"isDeleted" column:"account_codes.is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt *time.Time `json:"deletedAt" column:"account_codes.deleted_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt time.Time  `json:"createdAt" column:"account_codes.created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time  `json:"updatedAt" column:"account_codes.updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *AccountCode) TableName() string {
	return "account_codes"
}

func (m *AccountCode) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("account_codes_name_uindex", "name", true, "is_deleted = FALSE")
}
