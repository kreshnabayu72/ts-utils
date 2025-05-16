package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserAttribute struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName  string             `json:"fullName"      bson:"fullName"`
	Key       int                `json:"key"           bson:"key"`
	IsDefault bool               `json:"isDefault"     bson:"isDefault"`
}

type UserTransactionRole struct {
	ID                        primitive.ObjectID        `json:"_id,omitempty"             bson:"_id,omitempty"`
	User                      UserAttribute             `json:"user" bson:"user"`
	TransactionGroupAttribute TransactionGroupAttribute `json:"transactionGroupAttribute" bson:"transactionGroupAttribute"`
	Transaction               TransactionType           `json:"transaction"               bson:"transaction"`
	RoleType                  TransactionRoleType       `json:"roleType"                  bson:"roleType"`
	ApprovalLevel             int                       `json:"approvalLevel"             bson:"approvalLevel"`
	ApprovalType              ApprovalType              `json:"approvalType"              bson:"approvalType"`
	Group                     *Group                    `json:"group,omitempty"           bson:"group,omitempty"`
	ManagedGroups             []primitive.ObjectID      `json:"managedGroups,omitempty"   bson:"managedGroups,omitempty"`
	IsActive                  bool                      `json:"isActive"                  bson:"isActive"`
	IsManagedByParent         bool                      `json:"isManagedByParent"         bson:"isManagedByParent"`
	IsTotalControl            bool                      `json:"isTotalControl"            bson:"isTotalControl"`
	IsReadOnly                bool                      `json:"isReadOnly"                bson:"isReadOnly"`
	CreatedAt                 time.Time                 `json:"createdAt"                 bson:"createdAt"`
	UpdatedAt                 time.Time                 `json:"updatedAt"                 bson:"updatedAt"`
}

var ValidTransactionNames = []string{
	"borrowingRole",
	"disposalRole",
	"transferRole",
	"auditRole",
	"maintenanceRole",
	"repairRole",
	"trackingRole",
	"missingRole",
	"damagedRole",
}
