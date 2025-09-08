package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type QrRfidPair struct {
	ID         int        `json:"_id"        column:"qr_rfid_pairs.id"          sql:"BIGINT PRIMARY KEY"`
	QrId       int        `json:"qrId"       column:"qr_rfid_pairs.qr_id"       sql:"BIGINT"`
	RfidId     int        `json:"rfidId"     column:"qr_rfid_pairs.rfid_id"     sql:"BIGINT"`
	ReportedAt *time.Time `json:"reportedAt" column:"qr_rfid_pairs.reported_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt  time.Time  `json:"createdAt"  column:"qr_rfid_pairs.created_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time  `json:"updatedAt"  column:"qr_rfid_pairs.updated_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
