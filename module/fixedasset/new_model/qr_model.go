package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Qr struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Code   string    `json:"code" column:"code" sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(code) > 0)"`
	Status TagStatus `json:"status" column:"status" sql:"tag_status_enum NOT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (m *Qr) TableName() string {
	return "qrs"
}

func (m *Qr) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}
