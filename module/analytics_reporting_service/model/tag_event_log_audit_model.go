package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagTotalObject struct {
	Total     int `json:"total" bson:"total"`
	Qr        int `json:"qr" bson:"qr"`
	Rfid      int `json:"rfid" bson:"rfid"`
	Nfc       int `json:"nfc" bson:"nfc"`
	RfidNfc   int `json:"rfidNfc" bson:"rfidNfc"`
	RfidQr    int `json:"rfidQr" bson:"rfidQr"`
	NfcQr     int `json:"nfcQr" bson:"nfcQr"`
	RfidNfcQr int `json:"rfidNfcQr" bson:"rfidNfcQr"`
}

type TagEventLogAudit struct {
	ID                           primitive.ObjectID `json:"_id" bson:"_id"`
	TagEventLogID                primitive.ObjectID `json:"tagEventLogId" bson:"tagEventLogId"`
	Found                        TagTotalObject     `json:"found" bson:"found"`
	NotFound                     TagTotalObject     `json:"notFound" bson:"notFound"`
	MarkNotFoundAsDamagedMissing bool               `json:"markNotFoundAsDamagedMissing" bson:"markNotFoundAsDamagedMissing"`
	CreatedAt                    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt                    time.Time          `json:"updatedAt" bson:"updatedAt"`
}
