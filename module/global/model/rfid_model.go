package model

import (
	"time"

	common "gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rfid struct {
	ID primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`

	Group *common.AttributeEmbedded `json:"group"        bson:"group"`
	Asset *primitive.ObjectID       `json:"asset"        bson:"asset"`

	Status string `json:"status"        bson:"status"`

	RfidCode      string `json:"rfidCode"        bson:"rfidCode"`
	RfidShortCode string `json:"rfidShortCode"        bson:"rfidShortCode"`
	NfcCode       string `json:"nfcCode"        bson:"nfcCode"`
	NfcShortCode  string `json:"nfcShortCode"        bson:"nfcShortCode"`
	QrCode        string `json:"qrCode"        bson:"qrCode"`
	QrShortCode   string `json:"qrShortCode"        bson:"qrShortCode"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}
