package model

import (
	"time"

	common "gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Nfc struct {
	ID primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`

	Group            *common.AttributeEmbedded `json:"group"        bson:"group"`
	FixedAssetGroup  *common.AttributeEmbedded `bson:"fixedAssetGroup,omitempty"            json:"fixedAssetGroup,omitempty"`
	SupplyAssetGroup *common.AttributeEmbedded `bson:"supplyAssetGroup,omitempty"            json:"supplyAssetGroup,omitempty"`
	Asset            *primitive.ObjectID       `json:"asset"        bson:"asset"`

	Status string `json:"status"        bson:"status"`
	Type   string `json:"type"        bson:"type"`
	Module string `json:"module"        bson:"module"`

	Rfid          *primitive.ObjectID `json:"rfid"        bson:"rfid"`
	RfidCode      string              `json:"rfidCode"        bson:"rfidCode"`
	RfidShortCode string              `json:"rfidShortCode"        bson:"rfidShortCode"`
	Code          string              `json:"code"        bson:"code"`
	ShortCode     string              `json:"shortCode"        bson:"shortCode"`
	Qr            *primitive.ObjectID `json:"qr"        bson:"qr"`
	QrCode        string              `json:"qrCode"        bson:"qrCode"`
	QrShortCode   string              `json:"qrShortCode"        bson:"qrShortCode"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}
