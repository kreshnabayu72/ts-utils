package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Rfid struct {
	ID              int         `json:"_id"             column:"rfids.id"                sql:"BIGINT PRIMARY KEY"`
	Code            string      `json:"code"            column:"rfids.code"              sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(code) > 0)"`
	ShortCode       string      `json:"shortCode"       column:"rfids.short_code"        sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(short_code) > 0)"`
	Status          TagStatus   `json:"status"          column:"rfids.status"            sql:"tag_status_enum NOT NULL"`
	LastAssetStatus AssetStatus `json:"lastAssetStatus" column:"rfids.asset_status"      sql:"asset_status_enum NULL"`
	HolderId        *int        `json:"holderId"        column:"rfids.holder_id"         sql:"BIGINT"`
	ToBeReturnedAt  *time.Time  `json:"toBeReturnedAt"  column:"rfids.to_be_returned_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	ReportedById    *int        `json:"reportedById"    column:"rfids.reported_by_id"    sql:"BIGINT"`
	ReportedAt      *time.Time  `json:"reportedAt"      column:"rfids.reported_at"       sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt       time.Time   `json:"createdAt"       column:"rfids.created_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt       time.Time   `json:"updatedAt"       column:"rfids.updated_at"        sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Rfid) TableName() string {
	return "rfids"
}

func (m *Rfid) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("rfid_code_uindex", "code", true, "")

	builder.AddForeignKey(
		"fk_rfid_holder_id",
		"holder_id",
		"users",
		"id",
		"SET NULL",
	)
}
