package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

type Stock struct {
	ID               primitive.ObjectID `json:"_id,omitempty"    bson:"_id,omitempty"`
	ItemName         primitive.ObjectID `json:"itemName"         bson:"itemName"`
	ItemSKU          primitive.ObjectID `json:"itemSKU"          bson:"itemSKU"`
	Brand            primitive.ObjectID `json:"brand"            bson:"brand"`
	Model            primitive.ObjectID `json:"model"            bson:"model"`
	Category         primitive.ObjectID `json:"category"         bson:"category"`
	Group            primitive.ObjectID `json:"group"            bson:"group"`
	TagType          model.TagType      `json:"tagType"          bson:"tagType"`
	Tag              model.Tag          `json:"tag"              bson:"tag"`
	NameWithSequence string             `json:"nameWithSequence" bson:"nameWithSequence"`
	Sequence         int                `json:"sequence"         bson:"sequence"`
	Status           string             `json:"status"           bson:"status"`
	Source           string             `json:"source"           bson:"source"`
	ReceptionDate    time.Time          `json:"receptionDate"    bson:"receptionDate"`
	ProductionDate   *time.Time         `json:"productionDate"   bson:"productionDate"`
	ExpiredDate      *time.Time         `json:"expiredDate"      bson:"expiredDate"`
	UpdatedAt        time.Time          `json:"updatedAt"        bson:"updatedAt"`
	CreatedAt        time.Time          `json:"createdAt"        bson:"createdAt"`
}
