package new_model

import (
	"time"
)

type HandheldReader struct {
	ID                 int          `json:"_id"                column:"handheld_readers.id"`
	GroupID            int          `json:"groupId"            column:"handheld_readers.group_id"`
	ManagerId          int          `json:"managerId"          column:"handheld_readers.manager_id"`
	QrCode             *string      `json:"qrCode"             column:"handheld_readers.qr_code"`
	Image              string       `json:"image"              column:"handheld_readers.image"`
	SerialNumber       string       `json:"serialNumber"       column:"handheld_readers.serial_number"`
	DeviceNameId       int          `json:"deviceNameId"       column:"handheld_readers.device_name_id"`
	DeviceNameName     string       `json:"deviceNameName"     column:"handheld_readers.device_name_name"`
	DeviceNameSequence int          `json:"deviceNameSequence" column:"handheld_readers.device_name_sequence"`
	Status             ReaderStatus `json:"status"             column:"handheld_readers.status"`
	LastReportReason   ReaderStatus `json:"lastReportReason"   column:"handheld_readers.last_report_reason"`
	LastReportAt       time.Time    `json:"lastReportAt"       column:"handheld_readers.last_report_at"`
	CreatedAt          time.Time    `json:"createdAt"          column:"handheld_readers.created_at"`
	UpdatedAt          time.Time    `json:"updatedAt"          column:"handheld_readers.updated_at"`
}
