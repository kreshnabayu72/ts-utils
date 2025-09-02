package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Qr struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Code      string `json:"code" column:"code" sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(code) > 0)"`
	ShortCode string `json:"shortCode" column:"short_code" sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(short_code) > 0)"`

	Status TagStatus `json:"status" column:"status" sql:"tag_status_enum NOT NULL"`

	ReportedAt *time.Time `json:"reportedAt" column:"reported_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt  time.Time  `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time  `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Qr) TableName() string {
	return "qrs"
}

func (m *Qr) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}
