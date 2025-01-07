package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID                   primitive.ObjectID `json:"_id,omitempty"    bson:"_id,omitempty"`
	Name                 string             `json:"name"                 bson:"name"`
	Desc                 *string            `json:"desc,omitempty"                   bson:"desc,omitempty"`
	Type                 RoleType           `json:"type"             bson:"type"`
	ManagePermission     ManagePermission   `json:"managePermission" bson:"managePermission"`
	IsActive             bool               `json:"isActive"         bson:"isActive"`
	IsDefault            bool               `json:"isDefault"            bson:"isDefault"`
	IsDeleted            bool               `json:"isDeleted"            bson:"isDeleted"`
	DeleteFromDatabaseAt *time.Time         `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	Key                  int                `json:"key"                  bson:"key"`
	CreatedAt            time.Time          `json:"createdAt"        bson:"createdAt"`
	UpdatedAt            time.Time          `json:"updatedAt"        bson:"updatedAt"`
}
