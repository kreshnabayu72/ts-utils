package model

import (
	"time"

	common "gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Qr struct {
	ID primitive.ObjectID `json:"_id,omitempty"        bson:"_id,omitempty"`

	Group            *common.AttributeEmbedded `json:"group"        bson:"group"`
	FixedAssetGroup  *common.AttributeEmbedded `bson:"fixedAssetGroup,omitempty"            json:"fixedAssetGroup,omitempty"`
	SupplyAssetGroup *common.AttributeEmbedded `bson:"supplyAssetGroup,omitempty"            json:"supplyAssetGroup,omitempty"`
	Asset            *primitive.ObjectID       `json:"asset"        bson:"asset"`

	Status string `json:"status"        bson:"status"`
	Type   string `json:"type"        bson:"type"`
	Module string `json:"module"        bson:"module"`

	Code      string `json:"code"        bson:"code"`
	ShortCode string `json:"shortCode"        bson:"shortCode"`

	Nfc           *primitive.ObjectID `json:"nfc"        bson:"nfc"`
	NfcCode       string              `json:"nfcCode"        bson:"nfcCode"`
	NfcShortCode  string              `json:"nfcShortCode"        bson:"nfcShortCode"`
	Rfid          *primitive.ObjectID `json:"rfid"        bson:"rfid"`
	RfidCode      string              `json:"rfidCode"        bson:"rfidCode"`
	RfidShortCode string              `json:"rfidShortCode"        bson:"rfidShortCode"`

	CreatedAt time.Time `json:"createdAt"            bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"            bson:"updatedAt"`
}
