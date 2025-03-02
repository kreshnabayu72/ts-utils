package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

type PlacedItem struct {
	ItemSKU  primitive.ObjectID `json:"itemSKU"  bson:"itemSKU"`
	Capacity int                `json:"capacity" bson:"capacity"`
}

type Group struct {
	ID                    primitive.ObjectID  `json:"_id,omitempty"         bson:"_id,omitempty"`
	Key                   int                 `json:"key"                   bson:"key"`
	Name                  string              `json:"name"                  bson:"name"`
	FullPath              string              `json:"fullPath"              bson:"fullPath"`
	AliasCode             *string             `json:"aliasCode"             bson:"aliasCode"`
	Parent                *primitive.ObjectID `json:"parent"                bson:"parent"`
	AdministrativeGroup   primitive.ObjectID  `json:"administrativeGroup"   bson:"administrativeGroup"`
	Level                 int                 `json:"level"                 bson:"level"`
	Items                 []PlacedItem        `json:"items"                 bson:"items"`
	TagType               model.TagType       `json:"tagType"               bson:"tagType"`
	Tag                   model.Tag           `json:"tag"                   bson:"tag"`
	IsAdministrativeGroup bool                `json:"isAdministrativeGroup" bson:"isAdministrativeGroup"`
	IsDisposalGroup       bool                `json:"isDisposalGroup"       bson:"isDisposalGroup"`
	Image                 *string             `json:"image"                 bson:"image"`
	CreatedAt             time.Time           `json:"createdAt"             bson:"createdAt"`
	UpdatedAt             time.Time           `json:"updatedAt"             bson:"updatedAt"`
}
