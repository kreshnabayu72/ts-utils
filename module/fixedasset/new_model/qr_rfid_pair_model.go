package new_model

import "time"

type QrRfidPair struct {
	ID int `json:"id"         column:"id"`

	RfidId int `json:"rfidId"               column:"rfid_id"`
	QrId   int `json:"qrId"               column:"qr_id"`

	CreatedAt time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"updated_at"`
}
