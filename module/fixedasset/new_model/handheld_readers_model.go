package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type HandheldReader struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	GroupID   int     `json:"groupId" column:"group_id" sql:"BIGINT"`
	ManagerId int     `json:"managerId" column:"manager_id" sql:"BIGINT"`
	QrCode    *string `json:"qrCode" column:"qr_code" sql:"VARCHAR DEFAULT NULL"`
	Image     string  `json:"image" column:"image" sql:"VARCHAR DEFAULT NULL"`

	SerialNumber string `json:"serialNumber" column:"serial_number" sql:"VARCHAR NOT NULL"`

	DeviceNameId       int    `json:"deviceNameId" column:"device_name_id" sql:"BIGINT NOT NULL"`
	DeviceNameName     string `json:"deviceNameName" column:"device_name_name" sql:"VARCHAR NOT NULL"`
	DeviceNameSequence int    `json:"deviceNameSequence" column:"device_name_sequence" sql:"INTEGER NOT NULL"`

	Status           ReaderStatus `json:"status" column:"status" sql:"reader_status_enum NOT NULL DEFAULT 'Available'"`
	LastReportReason ReaderStatus `json:"lastReportReason" column:"last_report_reason" sql:"reader_status_enum DEFAULT NULL"`

	LastReportAt time.Time `json:"lastReportAt" column:"last_report_at" sql:"TIMESTAMPTZ DEFAULT NULL"`
	CreatedAt    time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt    time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (h *HandheldReader) TableName() string {
	return "handheld_readers"
}

func (h *HandheldReader) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(h)

	builder.AddForeignKey(
		"fk_handheld_readers_group_id",
		"group_id",
		"groups",
		"id",
		"SET NULL",
	)

	builder.AddForeignKey(
		"fk_handheld_readers_manager_id",
		"manager_id",
		"users",
		"id",
		"SET NULL",
	)

	builder.AddIndex("handheld_reader_qr_code_uindex", "qr_code", true, "")
	builder.AddIndex("handheld_reader_serial_number_uindex", "serial_number", true, "")
}
