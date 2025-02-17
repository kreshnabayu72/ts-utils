package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlacedItem struct {
	ItemSKU  primitive.ObjectID `json:"itemSKU"  bson:"itemSKU"`
	Capacity int                `json:"capacity" bson:"capacity"`
}

type Group struct {
	ID                    primitive.ObjectID `json:"_id,omitempty"         bson:"_id,omitempty"`
	Key                   int                `json:"key"                   bson:"key"`
	Name                  string             `json:"name"                  bson:"name"`
	FullPath              string             `json:"fullPath"              bson:"fullPath"`
	AliasCode             *string            `json:"aliasCode"             bson:"aliasCode"`
	Items                 []PlacedItem       `json:"items"                 bson:"items"`
	TagType               TagType            `json:"tagType"               bson:"tagType"`
	IsAdministrativeGroup bool               `json:"isAdministrativeGroup" bson:"isAdministrativeGroup"`
	IsDisposalGroup       bool               `json:"isDisposalGroup"       bson:"isDisposalGroup"`
	Tag                   Tag                `json:"tag"                   bson:"tag"`
	Image                 *string            `json:"image"                 bson:"image"`
	CreatedAt             time.Time          `json:"createdAt"             bson:"createdAt"`
	UpdatedAt             time.Time          `json:"updatedAt"             bson:"updatedAt"`
}
