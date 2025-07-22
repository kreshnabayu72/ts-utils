package new_model

import "time"

type HandheldReader struct {
	ID int `json:"_id" column:"id"`

	GroupID      int     `json:"groupId" column:"group_id"`
	ManagerId    int     `json:"managerId" column:"manager_id"`
	QrCode       *string `json:"qrCode" column:"qr_code"`
	Image        string  `json:"image" column:"image"`
	SerialNumber string  `json:"serialNumber" column:"serial_number"`

	DeviceNameId       int `json:"deviceNameId" column:"device_name_id"`
	DeviceNameName     int `json:"deviceNameName" column:"device_name_name"`
	DeviceNameSequence int `json:"deviceNameSequence" column:"device_name_sequence"`

	Status           ReaderStatus `json:"status" column:"status"`
	LastReportReason ReaderStatus `json:"lastReportReason" column:"last_report_reason"`

	LastReportAt time.Time `json:"lastReportAt" column:"last_report_at"`
	CreatedAt    time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" column:"updated_at"`
}
