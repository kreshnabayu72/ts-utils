package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IotActivityLog struct {
	ID            primitive.ObjectID  `json:"_id"         bson:"_id"`
	Reader        int                 `bson:"reader"                  json:"reader"`
	DeviceName    model.NameEmbedded  `bson:"deviceName"              json:"deviceName"`
	Group         *string             `bson:"group,omitempty"         json:"group,omitempty"`
	Activity      string              `bson:"activity"                json:"activity"`
	TransactionID *string             `bson:"transactionId,omitempty" json:"transactionId,omitempty"`
	AssetName     *model.NameEmbedded `bson:"assetName,omitempty"     json:"assetName,omitempty"`
	Manager       string              `bson:"manager"                 json:"manager"`
	AuditType     *string             `bson:"auditType,omitempty"     json:"auditType,omitempty"`
	PortNumber    *int                `bson:"portNumber,omitempty"    json:"portNumber,omitempty"`
	Type          string              `bson:"type"                    json:"type"`
	CreatedAt     time.Time           `bson:"createdAt"               json:"createdAt"`
	UpdatedAt     time.Time           `bson:"updatedAt"               json:"updatedAt"`
}
