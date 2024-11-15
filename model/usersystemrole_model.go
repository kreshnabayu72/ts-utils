package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSystemRole struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Role             AttributeEmbedded  `json:"role" bson:"role"`
	RoleName         RoleName           `json:"roleName" bson:"roleName"`
	Type             RoleType           `json:"type" bson:"type"`
	ManagePermission ManagePermission   `json:"managePermission" bson:"managePermission"`
	User             UserEmbedded       `json:"user" bson:"user"`
	Group            AttributeEmbedded  `json:"group" bson:"group"`
	IsActive         bool               `json:"isActive" bson:"isActive"`
	IsReadOnly       bool               `json:"isReadOnly" bson:"isReadOnly"`
	IsTotalControl   bool               `json:"isTotalControl" bson:"isTotalControl"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
}
