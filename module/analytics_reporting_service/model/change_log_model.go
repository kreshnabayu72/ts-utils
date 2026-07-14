// Package model defines the analytics-reporting-service database row types.
package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChangeLog struct {
	ID                     primitive.ObjectID `json:"_id" bson:"_id"`
	Context                string             `json:"context" bson:"context"`
	Action                 string             `json:"action" bson:"action"`
	ObjectID               string             `json:"objectId" bson:"objectId"`
	ObjectName             string             `json:"objectName" bson:"objectName"`
	ObjectNameWithSequence *string            `json:"objectNameWithSequence" bson:"objectNameWithSequence"`
	Field                  string             `json:"field" bson:"field"`
	OldValue               *string            `json:"oldValue" bson:"oldValue"`
	NewValue               *string            `json:"newValue" bson:"newValue"`
	UserID                 string             `json:"userId" bson:"userId"`
	UserFullName           string             `json:"userFullName" bson:"userFullName"`
	GroupID                *string            `json:"groupId" bson:"groupId"`
	SystemRoleKeys         []string           `json:"systemRoleKeys" bson:"systemRoleKeys"`
	TransactionRoleKeys    []string           `json:"transactionRoleKeys" bson:"transactionRoleKeys"`
	Module                 enum.ModuleAccess  `json:"module" bson:"module"`
	CreatedAt              time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt              time.Time          `json:"updatedAt" bson:"updatedAt"`
}
