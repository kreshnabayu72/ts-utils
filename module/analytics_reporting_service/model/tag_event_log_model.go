package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagEventLog struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	TotalTagScanned  *int               `json:"totalTagScanned" bson:"totalTagScanned"`
	Action           *string            `json:"action" bson:"action"`
	Module           *enum.ModuleAccess `json:"module" bson:"module"`
	Holder           *string            `json:"holder" bson:"holder"`
	ScannedRfid      *string            `json:"scannedRfid" bson:"scannedRfid"`
	ScannedNfc       *string            `json:"scannedNfc" bson:"scannedNfc"`
	TotalCombinedTag *int               `json:"totalCombinedTag" bson:"totalCombinedTag"`
	ModifiedBy       string             `json:"modifiedBy" bson:"modifiedBy"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
}
