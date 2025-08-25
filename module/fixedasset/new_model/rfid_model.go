package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Rfid struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Code            string      `json:"code" column:"code" sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(code) > 0)"`
	ShortCode       string      `json:"shortCode" column:"short_code" sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(short_code) > 0)"`
	Status          TagStatus   `json:"status" column:"status" sql:"tag_status_enum NOT NULL"`
	LastAssetStatus AssetStatus `json:"lastAssetStatus" column:"asset_status" sql:"asset_status_enum NULL"`

	HolderId *int `json:"holderId" column:"holder_id" sql:"BIGINT"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
