package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

type Attachment struct {
	ID        primitive.ObjectID `json:"_id"        bson:"_id"`
	Asset     primitive.ObjectID `json:"asset"      bson:"asset"`
	Type      string             `json:"type"       bson:"type"` // Possible values: "Purchase" or "Product"
	Name      string             `json:"name"       bson:"name"`
	File      string             `json:"file"       bson:"file"`
	FileName  string             `json:"fileName"       bson:"fileName"`
	CreatedAt time.Time          `json:"createdAt"  bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"  bson:"updatedAt"`
	CreatedBy model.UserEmbedded `json:"createdBy"  bson:"createdBy"`
}
