package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSystemRole struct {
	ID               primitive.ObjectID       `json:"_id,omitempty"    bson:"_id,omitempty"`
	RoleName         RoleName                 `json:"roleName"         bson:"roleName"`
	Type             RoleType                 `json:"type"             bson:"type"`
	RoleType         *string                  `json:"roleType,omitempty"             bson:"roleType,omitempty"` //"Global"|"Group"
	ManagePermission Permission               `json:"managePermission" bson:"managePermission"`
	User             model.UserEmbedded       `json:"user"             bson:"user"`
	Group            *model.AttributeEmbedded `json:"group"            bson:"group"`
	IsActive         bool                     `json:"isActive"         bson:"isActive"`
	IsReadOnly       bool                     `json:"isReadOnly"       bson:"isReadOnly"`
	IsTotalControl   bool                     `json:"isTotalControl"   bson:"isTotalControl"`
	CreatedAt        time.Time                `json:"createdAt"        bson:"createdAt"`
	UpdatedAt        time.Time                `json:"updatedAt"        bson:"updatedAt"`
}
