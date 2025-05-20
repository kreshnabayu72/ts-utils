package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

// PlacedItem represents the possible item sku that stock can be placed in a group, along with its capacity
type PlacedItem struct {
	ItemSKU        primitive.ObjectID `json:"itemSKU"        bson:"itemSKU"`
	Capacity       int                `json:"capacity"       bson:"capacity"`
	TotalStock     int                `json:"totalStock"     bson:"totalStock"`
	UnplacedStock  int                `json:"unplacedStock"  bson:"unplacedStock"`
	AvailableStock int                `json:"availableStock" bson:"availableStock"`
}

// Group represents the schema for individual group data
//
// It's meant to group other data to be used in transactions, presentations, etc.
type Group struct {
	ID                    primitive.ObjectID  `json:"_id,omitempty"         bson:"_id,omitempty"`
	Key                   int                 `json:"key"                   bson:"key"`
	Name                  string              `json:"name"                  bson:"name"`
	FullPath              string              `json:"fullPath"              bson:"fullPath"`
	AliasCode             *string             `json:"aliasCode"             bson:"aliasCode"`
	Parent                *primitive.ObjectID `json:"parent"                bson:"parent"`
	GroupType             string              `json:"groupType"             bson:"groupType"`
	CanPlaceAllItems      *bool				  `json:"canPlaceAllItems"      bson:"canPlaceAllItems"`
	AdministrativeGroup   primitive.ObjectID  `json:"administrativeGroup"   bson:"administrativeGroup"`
	Level                 int                 `json:"level"                 bson:"level"`
	Items                 []PlacedItem        `json:"items"                 bson:"items"`
	TagType               *model.TagType      `json:"tagType"               bson:"tagType"`
	Tag                   model.Tag           `json:"tag"                   bson:"tag"`
	IsAdministrativeGroup bool                `json:"isAdministrativeGroup" bson:"isAdministrativeGroup"`
	IsDisposalGroup       bool                `json:"isDisposalGroup"       bson:"isDisposalGroup"`
	Image                 *string             `json:"image"                 bson:"image"`
	IsDeleted             bool                `json:"isDeleted"             bson:"isDeleted"`
	CreatedAt             time.Time           `json:"createdAt"             bson:"createdAt"`
	UpdatedAt             time.Time           `json:"updatedAt"             bson:"updatedAt"`
}
