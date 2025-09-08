package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Qr struct {
	ID        int       `json:"_id"       column:"qrs.id"         sql:"BIGINT PRIMARY KEY"`
	Code      string    `json:"code"      column:"qrs.code"       sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(code) > 0)"`
	ShortCode string    `json:"shortCode" column:"qrs.short_code" sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(short_code) > 0)"`
	Status    TagStatus `json:"status"    column:"qrs.status"     sql:"tag_status_enum NOT NULL"`
	CreatedAt time.Time `json:"createdAt" column:"qrs.created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"qrs.updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Qr) TableName() string {
	return "qrs"
}

func (m *Qr) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}
