package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type QrRfidPair struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	QrId   int `json:"qrId"   column:"qr_id" sql:"BIGINT"`
	RfidId int `json:"rfidId" column:"rfid_id" sql:"BIGINT"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *QrRfidPair) TableName() string {
	return "qr_rfid_pairs"
}

func (m *QrRfidPair) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_qr_rfid_pair_qr_id",
		"qr_id",
		"qrs",
		"id",
		"SET NULL",
	)
	builder.AddForeignKey(
		"fk_qr_rfid_pair_rfid_id",
		"rfid_id",
		"rfids",
		"id",
		"SET NULL",
	)
}
